import pytest

from configs import *
from contracts import *

def test_getRedeemedTokenIds(contracts, sequence_length):
    iotx_stake = contracts[3]

    # At the beginning, there should be no redeemed token ID, even if 'deposit' has been called.
    iotx_stake.deposit()  # Todo:
    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    assert token_ids.len == 0

    # Once 'redeem' is called, the value of 'redeemedTokenCount' should increase correspondingly.
    # As a result, we will be able to query its ID.
    iotx_stake.redeem()  # Todo:
    assert iotx_stake.redeemedTokenCount() == 1
    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    assert token_ids.len == 1
    assert token_ids[0] == iotx_stake.tokenQueues(sequence_length-1, 0)

    # An empty result should be returned if inappropriate index values for i and j are passed
    # Todo
