def test_exchangeRatio(fn_isolation, contracts, users, delegates, oracles, admin, deadline):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # For the first 'deposit' request, the exchange ratio shouldn't change.
    default_exchange_ratio = 1000000000000000000
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.currentReserve() == amt
    assert uni_iotx.totalSupply() == amt
    assert iotx_staking.exchangeRatio() == default_exchange_ratio

    # The exchange ratio should be adjusted whenever any rewards are distributed from delegates
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    manager_fee = iotx_staking.getManagerFeeShares() * amt_reward / 1000
    user_reward = amt_reward - manager_fee
    expected_exchange_ratio1 = (amt + user_reward) * 1e18 / amt
    assert iotx_staking.getPendingReward() == amt_reward
    assert iotx_staking.currentReserve() == amt + user_reward
    assert uni_iotx.totalSupply() == amt
    assert iotx_staking.exchangeRatio() == expected_exchange_ratio1

    iotx_staking.updateReward({'from': oracles[0]})
    assert iotx_staking.getPendingReward() == 0
    assert iotx_staking.currentReserve() == amt + user_reward
    assert uni_iotx.totalSupply() == amt
    assert iotx_staking.exchangeRatio() == expected_exchange_ratio1

    # If a manager makes a 'withdrawManagerFee' request, the exchange ratio should theoretically maintain its previous value.
    # However, the actual value might vary from the theoretical one due to inherent errors in computational accuracy.
    iotx_staking.withdrawManagerFee(manager_fee, admin, {'from': admin})
    assert iotx_staking.getPendingReward() == 0
    assert iotx_staking.currentReserve() == amt + amt_reward
    assert uni_iotx.totalSupply() == amt + 9
    assert iotx_staking.exchangeRatio() == 1009091817364372065

    # If users make a 'redeem' request, the exchange ratio should theoretically maintain its previous value.
    # However, the actual value might vary from the theoretical one due to inherent errors in computational accuracy.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_staking.getPendingReward() == 0
    assert iotx_staking.currentReserve() == amt_reward
    assert uni_iotx.totalSupply() == 100
    assert iotx_staking.exchangeRatio() == default_exchange_ratio

    # If users make a 'deposit' request, the exchange ratio should theoretically maintain its previous value.
    # However, the actual value might vary from the theoretical one due to inherent errors in computational accuracy.
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.getPendingReward() == 0
    assert iotx_staking.currentReserve() == amt_reward + amt
    assert uni_iotx.totalSupply() == 100 + amt
    assert iotx_staking.exchangeRatio() == default_exchange_ratio
