import brownie
import pytest


def test_payDebts(fn_isolation, w3, contracts, users, delegates, oracle, admin, stake_amounts):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # The user's reward should be updated prior to debt payment.
    # A substantial debt can be paid off in several installments
    # until it is fully settled and its corresponding entry is removed.
    # Following the debt payment, both the 'accountedBalance' and the user's principal record should be adjusted.
    # The withdrawal of NFT during debt payment should not compromise the accuracy of reward recording.

    # Users deposit and redeem assets
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_staking.redeemAmountBase()
    cnt = 10
    amt_total = amt * cnt
    for i in range(0, cnt):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})

    uni_iotx.approve(iotx_staking, amt_total, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt_total, deadline, {'from': users[0], 'allow_revert': True})

    # Mock reward
    mock_reward_incr1 = 1000
    delegates[0].transfer(iotx_clear, mock_reward_incr1)

    # Oracle makes the first debt payment
    token_ids = iotx_staking.getRedeemedTokenIds(0, cnt)
    iotx_clear.unstake(token_ids[0:5], {'from': oracle, 'allow_revert': True})
    reward_rate1 = mock_reward_incr1 * 1e18 / amt_total
    reward_user01 = (reward_rate1 - 0) * amt_total / 1e18
    tx = iotx_clear.payDebts(token_ids[0:5], {'from': oracle, 'allow_revert': True})
    assert "Withdrawal" in tx.events
    assert "DebtPaid" in tx.events
    assert iotx_clear.getPaidDebtItemCount() == 0
    assert iotx_clear.totalDebts() == amt_total/2
    assert iotx_clear.accountedBalance() == mock_reward_incr1 + amt_total/2
    assert iotx_clear.balance() == iotx_clear.accountedBalance()

    user_info = iotx_clear.getUserInfo(users[0])
    debt_item1 = iotx_clear.getUnpaidDebtItem(1)
    assert user_info[0] == amt_total/2
    assert user_info[1] == amt_total/2
    assert user_info[2] == reward_user01
    assert user_info[3] == reward_rate1
    assert iotx_clear.rewardRate() == reward_rate1
    assert debt_item1[0] == users[0]
    assert debt_item1[1] == amt_total/2

    # Oracle makes the second payment
    iotx_clear.unstake(token_ids[5:cnt], {'from': oracle, 'allow_revert': True})
    tx = iotx_clear.payDebts(token_ids[5:cnt], {'from': oracle, 'allow_revert': True})
    assert "Withdrawal" in tx.events
    assert "DebtPaid" in tx.events
    assert iotx_clear.getPaidDebtItemCount() == 1
    assert iotx_clear.totalDebts() == 0
    assert iotx_clear.accountedBalance() == mock_reward_incr1 + amt_total
    assert iotx_clear.balance() == iotx_clear.accountedBalance()

    user_info = iotx_clear.getUserInfo(users[0])
    debt_item1 = iotx_clear.getUnpaidDebtItem(1)
    assert user_info[0] == 0
    assert user_info[1] == amt_total
    assert user_info[2] == reward_user01
    assert user_info[3] == reward_rate1
    assert iotx_clear.rewardRate() == reward_rate1
    assert debt_item1[0] == "0x0000000000000000000000000000000000000000"
    assert debt_item1[1] == 0

    # ---Revert path testing---

    # When the contract is on pause, the 'payDebts' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})
    iotx_clear.unpause({'from': admin})

    # Only the role of oracle has call permission.
    with brownie .reverts():
        iotx_clear.payDebts(token_ids, {'from': users[0], 'allow_revert': True})

    # Only accept debt tokens whose value is equal to the 'debtAmountBase' value
    iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[0], 'allow_revert': True})
    with brownie .reverts("USR007"):
        iotx_clear.payDebts([iotx_staking.getTokenId(0, 0)], {'from': oracle, 'allow_revert': True})

    iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[1], 'allow_revert': True})
    with brownie .reverts("USR007"):
        iotx_clear.payDebts([iotx_staking.getTokenId(1, 0)], {'from': oracle, 'allow_revert': True})

    # Passing an empty array of token IDs is not allowed
    iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[0], 'allow_revert': True})
    with brownie .reverts("USR008"):
        iotx_clear.payDebts([], {'from': oracle, 'allow_revert': True})

    # The total value of token IDs must not surpass the total value of the existing debt
    iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[2], 'allow_revert': True})
    token_id = iotx_staking.getTokenId(2, 10)
    with brownie .reverts("USR008"):
        iotx_clear.payDebts([token_id], {'from': oracle, 'allow_revert': True})

    # The token ID for payment will not be reused; therefore it must exist.
    with brownie .reverts("ERC721: invalid token ID"):
        iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})

    # The corresponding bucket(s) must be unstaked in advance.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("not an unstaked bucket"):
        iotx_clear.payDebts([token_id], {'from': oracle, 'allow_revert': True})

    # Note: in IoTeX testnet and mainnet, it can revert "not ready to withdraw".
    # https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L279
