def test_getStakedTokenCount(fn_isolation, w3, contracts, stake_amounts, users):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # Deposit any supported amount for each bucket,
    # The number of staked tokens should increase correspondingly, which can be queried from an external source.
    sequence_length = iotx_staking.sequenceLength()
    deadline = w3.eth.get_block('latest').timestamp+60
    for index in range(0, sequence_length):
        amt = stake_amounts[index]
        iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
        assert iotx_staking.getStakedTokenCount(index) == 1

    # The queried number of staked tokens should also reflect what has changed in case of merge or redeem.
    # Merge case
    amt = stake_amounts[0]
    for index in range(0, iotx_staking.commonRatio() - 1):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_staking.getStakedTokenCount(0) == 0
    assert iotx_staking.getStakedTokenCount(1) == 2

    # Redeem case
    amt = iotx_staking.redeemAmountBase()
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_staking.getStakedTokenCount(sequence_length - 1) == 0
