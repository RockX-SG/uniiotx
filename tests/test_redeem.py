import brownie
import pytest

from configs import *
from contracts import *


def test_redeem(w3, contracts, users, delegates, oracle):
    system_staking, uni_iotx, iotx_clear, iotx_staking = contracts[0], contracts[1], contracts[2], contracts[3]
    uint256_max = 115792089237316195423570985008687907853269984665640564039457584007913129639935

    # ---Happy path testing---

    # The exchange ratio will remain constant until the rewards are updated.
    # Initially, the amount of uniIOTX to burn is equivalent to the amount of IOTX to redeem.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_staking.redeemAmountBase()
    for i in range(0, 2):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    # The value transferred here will be considered as rewards from the delegate.
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    tx = iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    token_id = iotx_staking.tokenQueues(2, 0)
    unlocked_amt, _, unlocked_at, _, _ = system_staking.bucketOf(token_id)
    debt = iotx_clear.iotxDebts(1)
    user_info = iotx_clear.userInfos(users[0])
    assert len(tx.events["Transfer"]) == 3
    assert len(tx.events["Unlocked"]) == 1
    assert len(tx.events["Redeemed"]) == 1
    assert len(tx.events["DebtQueued"]) == 1
    assert unlocked_amt == amt
    assert unlocked_at != uint256_max
    assert iotx_staking.redeemedTokenCount() == 1
    assert iotx_staking.totalStaked() == amt
    assert uni_iotx.totalSupply() == amt
    assert uni_iotx.balanceOf(users[0]) == amt
    assert system_staking.balanceOf(iotx_staking) == 1
    assert system_staking.ownerOf(token_id) == iotx_clear
    assert system_staking.balanceOf(iotx_clear) == 1
    assert iotx_clear.totalDebts() == amt
    assert debt[0] == users[0]
    assert debt[1] == amt
    assert user_info[0] == amt

    # Regular updates to rewards can impact the exchange ratio's value.
    # The change in the exchange ratio should be taken into account.
    iotx_staking.updateReward({'from': oracle})
    manager_fee = iotx_staking.managerFeeShares() * amt_reward / 1000
    exchange_ratio1 = iotx_staking.exchangeRatio()
    assert exchange_ratio1 > 1e18
    max_to_burn = amt * 1e18 / exchange_ratio1
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    tx = iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    exchange_ratio2 = iotx_staking.exchangeRatio()
    assert exchange_ratio2 < exchange_ratio1
    token_id = iotx_staking.tokenQueues(2, 1)
    unlocked_amt, _, unlocked_at, _, _ = system_staking.bucketOf(token_id)
    debt = iotx_clear.iotxDebts(2)
    user_info = iotx_clear.userInfos(users[0])
    assert len(tx.events["Transfer"]) == 3
    assert len(tx.events["Unlocked"]) == 1
    assert len(tx.events["Redeemed"]) == 1
    assert len(tx.events["DebtQueued"]) == 1
    assert unlocked_amt == amt
    assert unlocked_at != uint256_max
    assert iotx_staking.redeemedTokenCount() == 2
    assert iotx_staking.totalPending() == amt_reward - manager_fee
    assert iotx_staking.totalStaked() == 0
    assert uni_iotx.totalSupply() == amt - max_to_burn
    assert uni_iotx.balanceOf(users[0]) == amt - max_to_burn
    assert system_staking.balanceOf(iotx_staking) == 0
    assert system_staking.ownerOf(token_id) == iotx_clear
    assert system_staking.balanceOf(iotx_clear) == 2
    assert iotx_clear.totalDebts() == amt*2
    assert debt[0] == users[0]
    assert debt[1] == amt
    assert user_info[0] == amt*2

    # ---Revert path testing---

    # The transaction of the redeem request should arrive within the deadline time.
    past_deadline = "1690514039"
    with brownie .reverts("Transaction expired"):
        iotx_staking.redeem(amt, past_deadline, {'from': users[0], 'allow_revert': True})

    # Users can can only redeem amounts that are multiples of the redeemAmountBase.
    with brownie .reverts("Invalid redeem amount"):
        iotx_staking.redeem(amt - 10, deadline, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("Invalid redeem amount"):
        iotx_staking.redeem(amt + 10, deadline, {'from': users[0], 'allow_revert': True})

    # Todo: Handle nonReentrant
    # Todo: Handle "ERC20: insufficient allowance"
