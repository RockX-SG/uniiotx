import brownie
import pytest

from configs import *
from contracts import *


def test_updateReward(contracts, users, delegates, oracle):
    iotx_stake = contracts[3]

    # ---Happy path testing---

    # No RewardUpdated event is emitted if delegates yield no rewards.
    tx = iotx_stake.updateReward({"from": oracle})
    assert "RewardUpdated" not in tx.events

    # The RewardUpdated event will be triggered if there is an existing reward that needs to be updated.
    # The value transferred here will be considered as rewards from the delegate.
    amt_reward = 100
    delegates[0].transfer(iotx_stake, amt_reward)
    tx = iotx_stake.updateReward({"from": oracle})
    manager_reward = iotx_stake.managerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert "RewardUpdated" in tx.events
    assert iotx_stake.accountedBalance() == amt_reward
    assert iotx_stake.accountedUserReward() == user_reward
    assert iotx_stake.accountedManagerReward() == manager_reward
    assert iotx_stake.totalPending() == user_reward

    # ---Revert path testing---

    # Only the role of Oracle has call permission.
    with brownie .reverts():
        iotx_stake.updateReward({"from": users[0]})
