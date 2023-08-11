import brownie


def test_withdrawManagerFee(fn_isolation, w3, contracts, users, delegates, admin, oracle):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    # An appropriate amount of uniIOTX should be minted, and the value of totalPending should be adjusted accordingly.
    # The value transferred here will be considered as rewards from the delegate.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_staking.redeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    amt_reward = 200
    delegates[0].transfer(iotx_staking, amt_reward)
    iotx_staking.updateReward({"from": oracle})
    manager_reward = iotx_staking.managerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_reward
    assert manager_reward == 20
    to_min = manager_reward / 2 * 1e18 / iotx_staking.exchangeRatio()
    tx = iotx_staking.withdrawManagerFee(manager_reward / 2, admin, {"from": admin})
    assert "ManagerFeeWithdrawed" in tx.events
    assert iotx_staking.accountedManagerReward() == manager_reward / 2
    assert iotx_staking.totalPending() == user_reward + manager_reward / 2
    assert uni_iotx.balanceOf(admin) == to_min

    # ---Revert path testing---

    # Ensure there is an adequate balance for the manager's fee withdrawal.
    with brownie .reverts("USR006"):
        iotx_staking.withdrawManagerFee(manager_reward, admin, {"from": admin})

    # Only the role of fee manager has call permission.
    with brownie .reverts():
        iotx_staking.withdrawManagerFee(manager_reward / 2, admin, {"from": users[0]})

    # Todo: Handle nonReentrant
