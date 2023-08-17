import brownie


def test_pause(fn_isolation, contracts, admin, oracles):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    tx = uni_iotx.pause({'from': iotx_staking, 'allow_revert': True})
    assert "Paused" in tx.events
    assert uni_iotx.paused()

    tx = iotx_clear.pause({'from': admin, 'allow_revert': True})
    assert "Paused" in tx.events
    assert iotx_clear.paused()

    tx = iotx_staking.pause({'from': admin, 'allow_revert': True})
    assert "Paused" in tx.events
    assert iotx_staking.paused()

    # ---Revert path testing---

    # Only the appropriate roles have call permissions.
    with brownie .reverts():
        uni_iotx.pause({'from': admin, 'allow_revert': True})
    with brownie .reverts():
        iotx_clear.pause({'from': oracles[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_staking.pause({'from': oracles[0], 'allow_revert': True})

    # It is callable only when not paused.
    with brownie .reverts("Pausable: paused"):
        uni_iotx.pause({'from': iotx_staking, 'allow_revert': True})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.pause({'from': admin, 'allow_revert': True})
    with brownie .reverts("Pausable: paused"):
        iotx_staking.pause({'from': admin, 'allow_revert': True})
