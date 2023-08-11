def test_getRedeemedTokenIds(fn_isolation, w3, contracts, users, deadline):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # At the beginning, there should be no redeemed token ID, even if 'deposit' has been called.
    sequence_length = iotx_staking.sequenceLength()
    amount = iotx_staking.redeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amount, 'allow_revert': True})
    token_ids = iotx_staking.getRedeemedTokenIds(0, 1)
    assert len(token_ids) == 0

    # Once 'redeem' is called, the value of 'redeemedTokenCount' should increase correspondingly.
    # As a result, we will be able to query its ID.
    uni_iotx.approve(iotx_staking, amount, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amount, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_staking.redeemedTokenCount() == 1
    token_ids = iotx_staking.getRedeemedTokenIds(0, 1)
    assert len(token_ids) == 1
    assert token_ids[0] == iotx_staking.getTokenId(sequence_length - 1, 0)

    # An empty result should be returned if inappropriate index values for i and j are passed
    # Todo: To optimize the implementation, we can use table-driven testing.
    token_ids = iotx_staking.getRedeemedTokenIds(0, 0)
    assert len(token_ids) == 0
    token_ids = iotx_staking.getRedeemedTokenIds(1, 0)
    assert len(token_ids) == 0
    token_ids = iotx_staking.getRedeemedTokenIds(0, iotx_staking.redeemedTokenCount() + 1)
    assert len(token_ids) == 0
