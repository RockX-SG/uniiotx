def test_currentReserve(fn_isolation, w3, contracts, users, delegates, oracle, admin):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # The current reserve is expected to increase once the 'deposit' call is successful.
    deadline = w3.eth.get_block('latest').timestamp+60
    redeem_amt = iotx_staking.redeemAmountBase()
    amt = redeem_amt + 10
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt

    # The rewards that have not been updated yet should not affect the current reserve.
    # The value sent here will be considered as rewards from the delegate
    amt_reward = 10
    delegates[0].transfer(iotx_staking, amt_reward)
    assert iotx_staking.currentReserve() == amt

    # The current reserve should be updated when rewards are updated.
    iotx_staking.updateReward({'from': oracle})
    manager_fee = amt_reward * iotx_staking.managerFeeShares() / 1000
    assert iotx_staking.currentReserve() == amt + amt_reward - manager_fee

    # If the manager fee is withdrawn, the current reserve should decrease accordingly.
    iotx_staking.withdrawManagerFee(manager_fee, admin, {'from': admin})
    assert iotx_staking.currentReserve() == amt + amt_reward

    # If users make a 'redeem' request, the current reserve should decrease accordingly.
    uni_iotx.approve(iotx_staking, redeem_amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(redeem_amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt + amt_reward - redeem_amt
