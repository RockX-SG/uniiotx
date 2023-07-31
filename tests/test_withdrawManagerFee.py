import brownie
import pytest

from configs import *
from contracts import *


def test_withdrawManagerFee(w3, contracts, users, delegates, admin, oracle):
    uni_iotx, iotx_stake = contracts[1], contracts[3]

    # ---Happy path testing---

    # An appropriate amount of uniIOTX should be minted, and the value of totalPending should be adjusted accordingly.
    # The value transferred here will be considered as rewards from the delegate.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()
    iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    amt_reward = 200
    delegates[0].transfer(iotx_stake, amt_reward)
    iotx_stake.updateReward({"from": oracle})
    manager_reward = iotx_stake.managerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert manager_reward == 20
    to_min = manager_reward/2 * iotx_stake.MULTIPLIER() / iotx_stake.exchangeRatio()
    tx = iotx_stake.withdrawManagerFee(manager_reward/2, admin, {"from": admin})
    assert "ManagerFeeWithdrawed" in tx.events
    assert iotx_stake.accountedManagerReward() == manager_reward/2
    assert iotx_stake.totalPending() == user_reward + manager_reward/2
    assert uni_iotx.balanceOf(admin) == to_min

    # ---Revert path testing---

    # Ensure there is an adequate balance for the manager's fee withdrawal.
    with brownie .reverts("Insufficient accounted manager reward"):
        iotx_stake.withdrawManagerFee(manager_reward, admin, {"from": admin})

    # Only the role of fee manager has call permission.
    with brownie .reverts():
        iotx_stake.withdrawManagerFee(manager_reward/2, admin, {"from": users[0]})

    # Todo: Handle nonReentrant
