import pytest

from configs import *
from contracts import *


def test_exchangeRatio(w3, contracts, users, delegates, oracle, admin):
    uni_iotx, iotx_stake = contracts[1], contracts[3]

    # The exchange ratio should be the same as the default value after the first 'deposit' process completes.
    default_exchange_ratio = 1000000000000000000
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()
    iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_stake.exchangeRatio() == default_exchange_ratio

    # The rewards that have not been updated yet should not affect the exchange ratio.
    # The value sent here will be considered as rewards from the delegate
    amt_reward = 100
    delegates[0].transfer(iotx_stake, amt_reward)
    assert iotx_stake.exchangeRatio() == default_exchange_ratio

    # Although the rewards have been received from the delegate, they have not been updated in our contract.
    # As a result, the exchange ratio will remain unchanged for the next 'deposit' request.
    iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    assert iotx_stake.exchangeRatio() == default_exchange_ratio

    # The exchange ratio should increase when rewards are updated in our contract.
    iotx_stake.updateReward({'from': oracle})
    assert iotx_stake.exchangeRatio() == 1004500000000000000

    # If the manager fee is withdrawn, the exchange ratio should increase accordingly.
    manager_fee = iotx_stake.managerFeeShares()*amt_reward / 1000
    iotx_stake.withdrawManagerFee(manager_fee, admin, {'from': admin})
    assert iotx_stake.exchangeRatio() == 1004547953420960567

    # If users make a 'redeem' request, the exchange ratio should decrease accordingly.
    uni_iotx.approve(iotx_stake, amt, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt, amt, deadline, {'from': users[0], 'allow_revert': True})
    assert iotx_stake.exchangeRatio() == 1004475385380407757
