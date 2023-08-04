import brownie
import pytest

from configs import *
from contracts import *

def test_claimPrincipals(w3, contracts, users, delegates, oracle, admin, stake_amounts):
    system_staking, uni_iotx, iotx_clear, iotx_stake = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # After repaying their debt, users' can claim their principal.
    # Users can claim any available principal amount at various times.
    # Once users have claimed their principals, the individual principal record
    # and the global 'accountedBalance' should be adjusted accordingly.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()

    iotx_stake.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    uni_iotx.approve(iotx_stake, amt, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})

    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    iotx_clear.unstake(token_ids, {'from': oracle, 'allow_revert': True})
    iotx_clear.payDebts(token_ids, {'from': oracle, 'allow_revert': True})

    accounted_balance0 = iotx_clear.accountedBalance()
    user0_balance0 = users[0].balance()
    user0_principal0 = iotx_clear.userInfos(users[0])[1]
    tx = iotx_clear.claimPrincipals(amt/2, users[0], {'from': users[0]})
    assert "PrincipalClaimed" in tx.events

    accounted_balance1 = iotx_clear.accountedBalance()
    user0_balance1 = users[0].balance()
    user0_principal1 = iotx_clear.userInfos(users[0])[1]

    assert accounted_balance1 == accounted_balance0 - amt / 2
    assert user0_balance1 == user0_balance0 + amt/2
    assert user0_principal1 == user0_principal0 - amt/2

    tx = iotx_clear.claimPrincipals(amt/2, users[0], {'from': users[0]})
    assert "PrincipalClaimed" in tx.events

    accounted_balance2 = iotx_clear.accountedBalance()
    user0_balance2 = users[0].balance()
    user0_principal2 = iotx_clear.userInfos(users[0])[1]

    assert accounted_balance2 == accounted_balance1 - amt / 2
    assert user0_balance2 == user0_balance1 + amt/2
    assert user0_principal2 == user0_principal1 - amt/2
    assert user0_principal2 == 0

    # ---Revert path testing---

    # When the contract is on pause, the 'claimPrincipals' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.claimPrincipals(amt/2, users[0], {'from': users[0]})
    iotx_clear.unpause({'from': admin})

    # Users should possess sufficient principals
    with brownie .reverts("Insufficient accounted principal"):
        iotx_clear.claimPrincipals(amt/2, users[0], {'from': users[0]})

    # Todo: Handle nonReentrant
