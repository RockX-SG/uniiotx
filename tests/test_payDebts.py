import brownie
import pytest

from configs import *
from contracts import *

def test_payDebts(w3, contracts, users, delegates, oracle, admin, stake_amounts):
    uni_iotx, iotx_clear, iotx_stake = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # The user's reward should be updated prior to debt payment.
    # Following the debt payment, both the 'accountedBalance' and the user's principal record should be adjusted.
    # The withdrawal of NFT during debt payment should not compromise the accuracy of reward recording.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()
    cnt = 10
    amt_total = amt * cnt
    for i in range(0, cnt):
        iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    uni_iotx.approve(iotx_stake, amt_total, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt_total, amt_total, deadline, {'from': users[0], 'allow_revert': True})
    token_ids = iotx_stake.getRedeemedTokenIds(0, cnt)
    mock_reward_incr1 = 1000
    delegates[0].transfer(iotx_clear, mock_reward_incr1)
    iotx_clear.unstake(token_ids, {'from': oracle, 'allow_revert': True})
    balance_user01 = users[0].balance()
    reward_rate1 = mock_reward_incr1 * 1e18 / amt_total
    reward_user01 = (reward_rate1 - 0) * amt_total / 1e18
    tx = iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})
    user_info = iotx_clear.userInfos(users[0])
    debt = iotx_clear.iotxDebts(1)
    assert "Withdrawal" in tx.events
    assert "DebtPaid" in tx.events
    assert users[0].balance() == balance_user01
    assert iotx_clear.headIndex() == 1
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 0
    assert iotx_clear.totalDebts() == 0
    assert iotx_clear.accountedBalance() == mock_reward_incr1 + amt_total
    assert iotx_clear.balance() == iotx_clear.accountedBalance()
    assert user_info[0] == 0
    assert user_info[1] == amt_total
    assert user_info[2] == reward_user01
    assert user_info[3] == reward_rate1
    assert iotx_clear.rewardRate() == reward_rate1
    assert debt[0] == "0x0000000000000000000000000000000000000000"
    assert debt[1] == 0

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
    iotx_stake.deposit(stake_amounts[0], deadline, {'from': users[0], 'value': stake_amounts[0], 'allow_revert': True})
    with brownie .reverts("Invalid token amount"):
        iotx_clear.payDebts([iotx_stake.tokenQueues(0, 0)], {'from': oracle, 'allow_revert': True})

    iotx_stake.deposit(stake_amounts[1], deadline, {'from': users[0], 'value': stake_amounts[1], 'allow_revert': True})
    with brownie .reverts("Invalid token amount"):
        iotx_clear.payDebts([iotx_stake.tokenQueues(1, 0)], {'from': oracle, 'allow_revert': True})

    # Passing an empty array of token IDs is not allowed
    iotx_stake.deposit(stake_amounts[0], deadline, {'from': users[0], 'value': stake_amounts[0], 'allow_revert': True})
    with brownie .reverts("Invalid total principal for debt payment"):
        iotx_clear.payDebts([], {'from': oracle, 'allow_revert': True})

    # The total value of token IDs must not surpass the total value of the existing debt
    iotx_stake.deposit(stake_amounts[2], deadline, {'from': users[0], 'value': stake_amounts[2], 'allow_revert': True})
    token_id = iotx_stake.tokenQueues(2, 10)
    with brownie .reverts("Invalid total principal for debt payment"):
        iotx_clear.payDebts([token_id], {'from': oracle, 'allow_revert': True})

    # The token ID for payment will not be reused; therefore it must exist.
    with brownie .reverts("ERC721: invalid token ID"):
        iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})

    # The corresponding bucket(s) must be unstaked in advance.
    uni_iotx.approve(iotx_stake, amt, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt, amt, deadline, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("not an unstaked bucket"):
        iotx_clear.payDebts([token_id], {'from': oracle, 'allow_revert': True})

    # Note: in IoTeX testnet and mainnet, it can revert "not ready to withdraw".
    # https://github.com/iotexproject/iip13-contracts/blob/main/src/SystemStaking.sol#L279
