import brownie


def test_withdrawManagerFee(fn_isolation, contracts, users, delegates, admin, oracles, deadline):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    # An appropriate amount of uniIOTX should be minted, and the value of totalPending should be adjusted accordingly.
    assert iotx_staking.exchangeRatio() == 1e18
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.getPendingReward() == 0
    amt_reward = 200
    delegates[0].transfer(iotx_staking, amt_reward)
    manager_reward = iotx_staking.getManagerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert iotx_staking.getPendingReward() == amt_reward
    assert iotx_staking.exchangeRatio() == 1018000000000000000
    assert manager_reward == 20
    to_min = manager_reward / 2 * 1e18 / iotx_staking.exchangeRatio()
    tx = iotx_staking.withdrawManagerFee(manager_reward / 2, admin, {"from": admin})
    assert "ManagerFeeWithdrawed" in tx.events
    assert iotx_staking.getPendingReward() == 0
    assert iotx_staking.getManagerReward() == manager_reward / 2
    assert iotx_staking.getTotalPending() == user_reward + manager_reward / 2
    assert uni_iotx.balanceOf(admin) == to_min
    assert iotx_staking.currentReserve() == 10190
    assert uni_iotx.totalSupply() == 10009
    assert iotx_staking.exchangeRatio() == 1018083724647816964

    # ---Revert path testing---

    # Ensure there is an adequate balance for the manager's fee withdrawal.
    with brownie .reverts("USR006"):
        iotx_staking.withdrawManagerFee(manager_reward, admin, {"from": admin})

    # Only the role of fee manager has call permission.
    with brownie .reverts():
        iotx_staking.withdrawManagerFee(manager_reward / 2, admin, {"from": users[0]})

    # Todo: Handle nonReentrant
