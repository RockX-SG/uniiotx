def test_currentReserve(fn_isolation, contracts, users, delegates, oracles, admin, deadline):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # The current reserve is expected to increase once the 'deposit' call is successful.
    redeem_amt = iotx_staking.getRedeemAmountBase()
    amt = redeem_amt + 10
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt

    # The current reserve should consistently reflect the most recently distributed reward
    amt_reward = 10
    delegates[0].transfer(iotx_staking, amt_reward)
    manager_fee = iotx_staking.getManagerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_fee
    assert iotx_staking.getPendingReward() == amt_reward
    assert iotx_staking.currentReserve() == amt + user_reward

    # The current reserve should be finally updated when rewards are updated.
    iotx_staking.updateReward({'from': oracles[0]})
    assert iotx_staking.currentReserve() == amt + user_reward

    # If the manager fee is withdrawn, the current reserve should increase accordingly.
    iotx_staking.withdrawManagerFee(manager_fee, admin, {'from': admin})
    assert iotx_staking.currentReserve() == amt + amt_reward

    # If users make a 'redeem' request, the current reserve should decrease accordingly.
    uni_iotx.approve(iotx_staking, redeem_amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(redeem_amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt + amt_reward - redeem_amt
