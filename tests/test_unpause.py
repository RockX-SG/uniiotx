import brownie


def test_unpause(fn_isolation, contracts, admin, oracles):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    uni_iotx.pause({'from': iotx_staking, 'allow_revert': True})
    tx = uni_iotx.unpause({'from': iotx_staking, 'allow_revert': True})
    assert "Unpaused" in tx.events
    assert not uni_iotx.paused()

    iotx_clear.pause({'from': admin, 'allow_revert': True})
    tx = iotx_clear.unpause({'from': admin, 'allow_revert': True})
    assert "Unpaused" in tx.events
    assert not iotx_clear.paused()

    iotx_staking.pause({'from': admin, 'allow_revert': True})
    tx = iotx_staking.unpause({'from': admin, 'allow_revert': True})
    assert "Unpaused" in tx.events
    assert not iotx_staking.paused()

    # ---Revert path testing---

    # Only the appropriate roles have call permissions.
    with brownie .reverts():
        uni_iotx.unpause({'from': admin, 'allow_revert': True})
    with brownie .reverts():
        iotx_clear.unpause({'from': oracles[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_staking.unpause({'from': oracles[0], 'allow_revert': True})

    # It is callable only when paused.
    with brownie .reverts("Pausable: not paused"):
        uni_iotx.unpause({'from': iotx_staking, 'allow_revert': True})
    with brownie .reverts("Pausable: not paused"):
        iotx_clear.unpause({'from': admin, 'allow_revert': True})
    with brownie .reverts("Pausable: not paused"):
        iotx_staking.unpause({'from': admin, 'allow_revert': True})
