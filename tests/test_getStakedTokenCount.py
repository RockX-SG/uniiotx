import pytest

from configs import *
from contracts import *

def test_getStakedTokenCount(w3, contracts, sequence_length, stake_amounts, common_ratio, users):
    uni_iotx, iotx_stake = contracts[1], contracts[3]

    # Deposit any supported amount for each bucket,
    # The number of staked tokens should increase correspondingly, which can be queried from an external source.
    deadline = w3.eth.get_block('latest').timestamp+60
    for index in range(0, sequence_length):
        amt = stake_amounts[index]
        iotx_stake.deposit(amt, deadline,
                           {'from': users[0], 'value': amt, 'gas_limit': '1000000000000000000', 'allow_revert': True})
        assert iotx_stake.getStakedTokenCount(index) == 1

    # The queried number of staked tokens should also reflect what has changed in case of merge or redeem.
    # Merge case
    amt = stake_amounts[0]
    for index in range(0, common_ratio-1):
        iotx_stake.deposit(amt, deadline,
                           {'from': users[0], 'value': amt, 'gas_limit': '1000000000000000000', 'allow_revert': True})
    assert iotx_stake.getStakedTokenCount(0) == 0
    assert iotx_stake.getStakedTokenCount(1) == 2

    # Redeem case
    amt = stake_amounts[sequence_length-1]
    uni_iotx.approve(iotx_stake, amt, {'from': users[0], 'gas_limit': '1000000000000000000', 'allow_revert': True})
    iotx_stake.redeem(amt, amt, deadline,
                      {'from': users[0], 'gas_limit': '1000000000000000000', 'allow_revert': True})
    assert iotx_stake.getStakedTokenCount(sequence_length-1) == 0
