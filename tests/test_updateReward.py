import brownie


def test_updateReward(fn_isolation, contracts, users, delegates, oracles):
    iotx_staking = contracts[3]

    # ---Happy path testing---

    # No RewardUpdated event is emitted if delegates yield no rewards.
    tx = iotx_staking.updateReward({"from": oracles[0]})
    assert "RewardUpdated" not in tx.events

    # The RewardUpdated event will be triggered if there is an existing reward that needs to be updated.
    # The value transferred here will be considered as rewards from the delegate.
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    tx = iotx_staking.updateReward({"from": oracles[0]})
    manager_reward = iotx_staking.getManagerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert "RewardUpdated" in tx.events
    assert iotx_staking.getAccountedBalance() == amt_reward
    assert iotx_staking.getUserReward() == user_reward
    assert iotx_staking.getManagerReward() == manager_reward
    assert iotx_staking.getTotalPending() == user_reward

    # ---Revert path testing---

    # Only the role of Oracle has call permission.
    with brownie .reverts():
        iotx_staking.updateReward({"from": users[0]})
