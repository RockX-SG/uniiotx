import brownie


def test_updateDelegates(fn_isolation, contracts, users, delegates, oracles, admin, deadline):
    system_staking, uni_iotx, iotx_clear, iotx_staking = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # The token ID owner should have the ability to update delegates.
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    token_id = iotx_staking.getTokenId(iotx_staking.sequenceLength() - 1, 0)
    assert system_staking.ownerOf(token_id) == iotx_staking
    tx = iotx_staking.updateDelegates([token_id], delegates[1], {'from': oracles[0]})
    assert "DelegatesUpdated" in tx.events

    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    token_ids = iotx_staking.getRedeemedTokenIdSlice(0, 1)
    assert system_staking.ownerOf(token_ids[0]) == iotx_clear
    tx = iotx_clear.updateDelegates(token_ids, delegates[0], {'from': oracles[0]})
    assert "DelegatesUpdated" in tx.events

    # ---Revert path testing---

    # When the contract is on pause, the 'updateDelegates' function will not operate.
    iotx_staking.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_staking.updateDelegates([token_id], delegates[1], {'from': oracles[0]})
    iotx_staking.unpause({'from': admin})

    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.updateDelegates(token_ids, delegates[0], {'from': oracles[0]})
    iotx_clear.unpause({'from': admin})

    # Only the role of oracle has call permission.
    with brownie .reverts():
        iotx_staking.updateDelegates([token_id], delegates[1], {'from': admin})

    with brownie .reverts():
        iotx_clear.updateDelegates([token_id], delegates[1], {'from': admin})

    # Unable to switch to the same delegate
    with brownie .reverts():
        iotx_staking.updateDelegates([token_id], delegates[0], {'from': admin})

    with brownie .reverts():
        iotx_clear.updateDelegates([token_id], delegates[0], {'from': admin})


