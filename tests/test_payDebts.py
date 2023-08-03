import brownie
import pytest

from configs import *
from contracts import *

def test_updateDelegates(w3, contracts, users, delegates, oracle, admin):
    system_staking, uni_iotx, iotx_clear, iotx_stake = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # The value transferred here will be considered as rewards from the delegate.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()
    cnt = 10
    amt_total = amt * cnt
    for i in range(0, cnt):
        iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    uni_iotx.approve(iotx_stake, amt_total, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt_total, amt_total, deadline, {'from': users[0], 'allow_revert': True})
    token_ids = iotx_stake.getRedeemedTokenIds(0, cnt)
    reward_incr1 = 1000
    delegates[0].transfer(iotx_clear, reward_incr1)
    iotx_clear.unstake(token_ids, {'from': oracle, 'allow_revert': True})
    balance_user01 = users[0].balance()
    reward_rate1 = reward_incr1 * 1e18 / amt_total
    reward_user01 = (reward_rate1 - 0) * amt_total / 1e18
    tx = iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})
    user_info = iotx_clear.userInfos(users[0])
    debt = iotx_clear.iotxDebts(1)
    assert "DebtPaid" in tx.events
    assert users[0].balance() == amt_total + balance_user01
    assert iotx_clear.headIndex() == 1
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 0
    assert iotx_clear.totalDebts() == 0
    assert iotx_clear.accountedBalance() == reward_incr1
    assert user_info[0] == 0
    assert user_info[1] == reward_user01
    assert user_info[2] == reward_rate1
    assert iotx_clear.rewardRate() == reward_rate1
    assert debt[0] == "0x0000000000000000000000000000000000000000"
    assert debt[1] == 0

    # Todo: add more testing case
    # Todo: to be optimize

    # ---Revert path testing---

    # # When the contract is on pause, the 'updateDelegates' function will not operate.
    # iotx_stake.pause({'from': admin})
    # with brownie .reverts("Pausable: paused"):
    #     iotx_stake.updateDelegates([token_id], delegates[1], {'from': oracle})
    # iotx_stake.unpause({'from': admin})
    #
    # iotx_clear.pause({'from': admin})
    # with brownie .reverts("Pausable: paused"):
    #     iotx_clear.updateDelegates(token_ids, delegates[0], {'from': oracle})
    # iotx_clear.unpause({'from': admin})
    #
    # # Only the role of oracle has call permission.
    # with brownie .reverts():
    #     iotx_stake.updateDelegates([token_id], delegates[1], {'from': admin})
    #
    # with brownie .reverts():
    #     iotx_clear.updateDelegates([token_id], delegates[1], {'from': admin})
    #
    # # Unable to switch to the same delegate
    # with brownie .reverts():
    #     iotx_stake.updateDelegates([token_id], delegates[0], {'from': admin})
    #
    # with brownie .reverts():
    #     iotx_clear.updateDelegates([token_id], delegates[0], {'from': admin})


