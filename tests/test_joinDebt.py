import brownie
import pytest

from configs import *
from contracts import *


def test_joinDebt(contracts, users, delegates, admin):
    iotx_clear, iotx_stake = contracts[2], contracts[3]

    # ---Happy path testing---

    # The debts of different users should be recorded/queued separately and appropriately.
    # The user's dynamic reward rate should be updated before a new debt record is eventually added.
    # The value transferred here will be considered as rewards from the delegate.

    account = users[0]
    amount = iotx_stake.redeemAmountBase()
    tx = iotx_clear.joinDebt(account, amount, {"from": iotx_stake})
    assert "DebtAdded" in tx.events
    assert iotx_clear.totalDebts() == amount
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 1
    assert iotx_clear.rearIndex() == 1
    assert iotx_clear.iotxDebts(1)[0] == account
    assert iotx_clear.iotxDebts(1)[1] == amount
    assert iotx_clear.userInfos(account)[0] == amount
    assert iotx_clear.rewardRate() == 0

    reward_incr1 = 100
    delegates[0].transfer(iotx_clear, reward_incr1)
    reward_rate1 = reward_incr1 * 1e18 / iotx_clear.totalDebts()
    account = users[1]
    tx = iotx_clear.joinDebt(account, amount, {"from": iotx_stake})
    assert "DebtAdded" in tx.events
    assert iotx_clear.totalDebts() == amount*2
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 2
    assert iotx_clear.rearIndex() == 2
    assert iotx_clear.iotxDebts(2)[0] == account
    assert iotx_clear.iotxDebts(2)[1] == amount
    assert iotx_clear.userInfos(account)[0] == amount
    assert iotx_clear.userInfos(account)[1] == 0
    assert iotx_clear.userInfos(account)[2] == reward_rate1
    assert iotx_clear.accountedBalance() == reward_incr1
    assert iotx_clear.rewardRate() == reward_rate1

    reward_incr2 = 200
    delegates[0].transfer(iotx_clear, reward_incr2)
    account = users[0]
    reward_rate2 = reward_rate1 + (reward_incr2 * 1e18 / iotx_clear.totalDebts())
    debt_user0 = iotx_clear.userInfos(account)[0]
    reward_rate_user0 = iotx_clear.userInfos(account)[2]
    reward_user0 = (reward_rate2 - reward_rate_user0) * debt_user0 / 1e18
    tx = iotx_clear.joinDebt(account, amount, {"from": iotx_stake})
    assert "DebtAdded" in tx.events
    assert iotx_clear.totalDebts() == amount*3
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 3
    assert iotx_clear.rearIndex() == 3
    assert iotx_clear.iotxDebts(1)[0] == account
    assert iotx_clear.iotxDebts(1)[1] == amount
    assert iotx_clear.iotxDebts(3)[0] == account
    assert iotx_clear.iotxDebts(3)[1] == amount
    assert iotx_clear.userInfos(account)[0] == amount*2
    assert iotx_clear.userInfos(account)[1] == reward_user0
    assert iotx_clear.userInfos(account)[2] == reward_rate2
    assert iotx_clear.accountedBalance() == reward_incr1 + reward_incr2
    assert iotx_clear.rewardRate() == reward_rate2

    reward_incr3 = 300
    delegates[0].transfer(iotx_clear, reward_incr3)
    account = users[1]
    reward_rate3 = reward_rate2 + (reward_incr3 * 1e18 / iotx_clear.totalDebts())
    debt_user1 = iotx_clear.userInfos(account)[0]
    reward_rate_user1 = iotx_clear.userInfos(account)[2]
    reward_user1 = (reward_rate3 - reward_rate_user1) * debt_user1 / 1e18
    tx = iotx_clear.joinDebt(account, amount, {"from": iotx_stake})
    assert "DebtAdded" in tx.events
    assert iotx_clear.totalDebts() == amount*4
    assert iotx_clear.rearIndex() - iotx_clear.headIndex() == 4
    assert iotx_clear.rearIndex() == 4
    assert iotx_clear.iotxDebts(2)[0] == account
    assert iotx_clear.iotxDebts(2)[1] == amount
    assert iotx_clear.iotxDebts(4)[0] == account
    assert iotx_clear.iotxDebts(4)[1] == amount
    assert iotx_clear.userInfos(account)[0] == amount*2
    assert iotx_clear.userInfos(account)[1] == reward_user1
    assert iotx_clear.userInfos(account)[2] == reward_rate3
    assert iotx_clear.accountedBalance() == reward_incr1 + reward_incr2 + reward_incr3
    assert iotx_clear.rewardRate() == reward_rate3

    # ---Revert path testing---

    # When the contract is on pause, the 'joinDebt' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.joinDebt(account, amount, {"from": iotx_stake})
    iotx_clear.unpause({'from': admin})

    # Only the role of stake has call permission.
    with brownie .reverts():
        iotx_clear.joinDebt(account, amount, {"from": admin})
