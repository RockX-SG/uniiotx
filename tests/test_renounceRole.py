import brownie


def test_renounceRole(fn_isolation, contracts, users, roles, admin):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # First time renouncing
    tx = uni_iotx.renounceRole(roles[0], iotx_staking, {'from': iotx_staking, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not uni_iotx.hasRole(roles[0], iotx_staking)

    tx = iotx_clear.renounceRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not iotx_clear.hasRole(roles[0], admin)

    tx = iotx_staking.revokeRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not iotx_staking.hasRole(roles[0], admin)

    # No repeated renouncing
    tx = uni_iotx.renounceRole(roles[0], iotx_staking, {'from': iotx_staking, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not uni_iotx.hasRole(roles[0], iotx_staking)

    tx = iotx_clear.renounceRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not iotx_clear.hasRole(roles[0], admin)

    tx = iotx_staking.renounceRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not iotx_staking.hasRole(roles[0], admin)

    # ---Revert path testing---

    # Can only renounce roles for self
    with brownie .reverts("AccessControl: can only renounce roles for self"):
        uni_iotx.renounceRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("AccessControl: can only renounce roles for self"):
        iotx_clear.renounceRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("AccessControl: can only renounce roles for self"):
        iotx_staking.renounceRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
