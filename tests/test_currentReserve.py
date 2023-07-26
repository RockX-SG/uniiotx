import pytest

from configs import *
from contracts import *


def test_currentReserve(w3, contracts, users, delegates, oracle, admin):
    uni_iotx, iotx_stake = contracts[1], contracts[3]

    # The current reserve is expected to increase once the 'deposit' call is successful.
    deadline = w3.eth.get_block('latest').timestamp+60
    redeem_amt = iotx_stake.redeemAmountBase()
    amt = redeem_amt + 10
    iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_stake.currentReserve() == amt

    # The rewards that have not been updated yet should not affect the current reserve.
    # The value sent here will be considered as rewards from the delegate
    amt_reward = 10
    delegates[0].transfer(iotx_stake, amt_reward)
    assert iotx_stake.currentReserve() == amt

    # The current reserve should be updated when rewards are updated.
    iotx_stake.updateReward({'from': oracle})
    manager_fee = amt_reward * iotx_stake.managerFeeShares()/1000
    assert iotx_stake.currentReserve() == amt + amt_reward - manager_fee

    # If the manager fee is withdrawn, the current reserve should decrease accordingly.
    iotx_stake.withdrawManagerFee(manager_fee, admin, {'from': admin})
    assert iotx_stake.currentReserve() == amt + amt_reward

    # If users make a 'redeem' request, the current reserve should decrease accordingly.
    uni_iotx.approve(iotx_stake, redeem_amt, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(redeem_amt, redeem_amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_stake.currentReserve() == amt + amt_reward - redeem_amt
