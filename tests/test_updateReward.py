import brownie
import pytest


def test_updateReward(fn_isolation, contracts, users, delegates, oracle):
    iotx_staking = contracts[3]

    # ---Happy path testing---

    # No RewardUpdated event is emitted if delegates yield no rewards.
    tx = iotx_staking.updateReward({"from": oracle})
    assert "RewardUpdated" not in tx.events

    # The RewardUpdated event will be triggered if there is an existing reward that needs to be updated.
    # The value transferred here will be considered as rewards from the delegate.
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    tx = iotx_staking.updateReward({"from": oracle})
    manager_reward = iotx_staking.managerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert "RewardUpdated" in tx.events
    assert iotx_staking.accountedBalance() == amt_reward
    assert iotx_staking.accountedUserReward() == user_reward
    assert iotx_staking.accountedManagerReward() == manager_reward
    assert iotx_staking.totalPending() == user_reward

    # ---Revert path testing---

    # Only the role of Oracle has call permission.
    with brownie .reverts():
        iotx_staking.updateReward({"from": users[0]})
