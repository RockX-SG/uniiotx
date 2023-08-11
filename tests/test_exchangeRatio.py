def test_exchangeRatio(fn_isolation, w3, contracts, users, delegates, oracle, admin):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # The exchange ratio should be the same as the default value after the first 'deposit' process completes.
    default_exchange_ratio = 1000000000000000000
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_staking.redeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt
    assert uni_iotx.totalSupply() == amt
    assert iotx_staking.exchangeRatio() == default_exchange_ratio

    # The rewards that have not been updated yet should not affect the exchange ratio.
    # The value sent here will be considered as rewards from the delegate
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    assert iotx_staking.currentReserve() == amt
    assert uni_iotx.totalSupply() == amt
    assert iotx_staking.exchangeRatio() == default_exchange_ratio

    # Although the rewards have been received from the delegate, they have not been updated in our contract.
    # As a result, the exchange ratio will remain unchanged for the next 'deposit' request.
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt * 2
    assert uni_iotx.totalSupply() == amt*2
    assert iotx_staking.exchangeRatio() == default_exchange_ratio

    # The exchange ratio should increase when rewards are updated in our contract.
    iotx_staking.updateReward({'from': oracle})
    manager_fee = iotx_staking.managerFeeShares() * amt_reward / 1000
    current_reserve1 = iotx_staking.currentReserve()
    total_supply1 = uni_iotx.totalSupply()
    exchange_ratio1 = iotx_staking.exchangeRatio()
    assert exchange_ratio1 > default_exchange_ratio
    assert current_reserve1 == amt*2 + (amt_reward-manager_fee)
    assert total_supply1 == amt * 2
    assert exchange_ratio1 == 1004500000000000000

    # If the manager fee is withdrawn, the exchange ratio should increase accordingly.
    iotx_staking.withdrawManagerFee(manager_fee, admin, {'from': admin})
    current_reserve2 = iotx_staking.currentReserve()
    total_supply2 = uni_iotx.totalSupply()
    exchange_ratio2 = iotx_staking.exchangeRatio()
    assert exchange_ratio2 > exchange_ratio1
    assert current_reserve2 == amt*2 + amt_reward
    assert total_supply2 == total_supply1 + 1e18 * manager_fee / exchange_ratio1
    assert exchange_ratio2 == 1004547953420960567

    # If users make a 'redeem' request, the exchange ratio should decrease accordingly.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    current_reserve3 = iotx_staking.currentReserve()
    total_supply3 = uni_iotx.totalSupply()
    exchange_ratio3 = iotx_staking.exchangeRatio()
    assert exchange_ratio3 < exchange_ratio2
    assert current_reserve3 == (amt*2+amt_reward) - amt
    assert total_supply3 == total_supply2 - 1e18 * amt / exchange_ratio2
    assert exchange_ratio3 == 1004475385380407757
