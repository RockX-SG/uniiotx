import brownie
import pytest

from configs import *
from contracts import *

def test_updateDelegates(w3, contracts, users, delegates, oracle, admin):
    system_staking, uni_iotx, iotx_clear, iotx_stake = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # The token ID owner should have the ability to update delegates.
    deadline = w3.eth.get_block('latest').timestamp+60
    amt = iotx_stake.redeemAmountBase()
    iotx_stake.deposit(amt, deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    token_id = iotx_stake.tokenQueues(iotx_stake.sequenceLength()-1, 0)
    assert system_staking.ownerOf(token_id) == iotx_stake
    iotx_stake.updateDelegates([token_id], delegates[1], {'from': oracle})

    uni_iotx.approve(iotx_stake, amt, {'from': users[0], 'allow_revert': True})
    iotx_stake.redeem(amt, amt, deadline, {'from': users[0], 'allow_revert': True})
    token_ids = iotx_stake.getRedeemedTokenIds(0, 1)
    assert system_staking.ownerOf(token_ids[0]) == iotx_clear
    iotx_clear.updateDelegates(token_ids, delegates[0], {'from': oracle})

    # ---Revert path testing---

    # When the contract is on pause, the 'updateDelegates' function will not operate.
    iotx_stake.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_stake.updateDelegates([token_id], delegates[1], {'from': oracle})
    iotx_stake.unpause({'from': admin})

    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.updateDelegates(token_ids, delegates[0], {'from': oracle})
    iotx_clear.unpause({'from': admin})

    # Only the role of oracle has call permission.
    with brownie .reverts():
        iotx_stake.updateDelegates([token_id], delegates[1], {'from': admin})

    with brownie .reverts():
        iotx_clear.updateDelegates([token_id], delegates[1], {'from': admin})

    # Unable to switch to the same delegate
    with brownie .reverts():
        iotx_stake.updateDelegates([token_id], delegates[0], {'from': admin})

    with brownie .reverts():
        iotx_clear.updateDelegates([token_id], delegates[0], {'from': admin})


