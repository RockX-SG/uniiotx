import pytest

from configs import *
from contracts import *

def test_getRedeemedTokenIds(w3, contracts, sequence_length, stake_amounts, users):
    uni_iotx, iotx_stake = contracts[1], contracts[3]

    # At the beginning, there should be no redeemed token ID, even if 'deposit' has been called.
    amount = stake_amounts[sequence_length-1]
    deadline = w3.eth.get_block('latest').timestamp+60
    iotx_stake.deposit(amount, deadline, {'from': users[0], 'value': amount, 'allow_revert': True})
    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    assert len(token_ids) == 0

    # Once 'redeem' is called, the value of 'redeemedTokenCount' should increase correspondingly.
    # As a result, we will be able to query its ID.
    uni_iotx.approve(iotx_stake, amount, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amount, amount, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_stake.redeemedTokenCount() == 1
    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    assert len(token_ids) == 1
    assert token_ids[0] == iotx_stake.tokenQueues(sequence_length-1, 0)

    # An empty result should be returned if inappropriate index values for i and j are passed
    # Todo: To optimize the implementation, we can use table-driven testing.
    token_ids = iotx_stake.getRedeemedTokenIds(0, 0)
    assert len(token_ids) == 0
    token_ids = iotx_stake.getRedeemedTokenIds(1, 0)
    assert len(token_ids) == 0
    token_ids = iotx_stake.getRedeemedTokenIds(0, iotx_stake.redeemedTokenCount()+1)
    assert len(token_ids) == 0