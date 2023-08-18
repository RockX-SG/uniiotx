import brownie


def test_revokeRole(fn_isolation, contracts, users, roles, admin):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # First time revoking
    tx = uni_iotx.revokeRole(roles[0], iotx_staking, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not uni_iotx.hasRole(roles[0], iotx_staking)

    tx = iotx_clear.revokeRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not iotx_clear.hasRole(roles[0], admin)

    tx = iotx_staking.revokeRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" in tx.events
    assert not iotx_staking.hasRole(roles[0], admin)

    # No repeated revoking
    tx = uni_iotx.revokeRole(roles[0], iotx_staking, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not uni_iotx.hasRole(roles[0], iotx_staking)

    tx = iotx_clear.revokeRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not iotx_clear.hasRole(roles[0], admin)

    tx = iotx_staking.revokeRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleRevoked" not in tx.events
    assert not iotx_staking.hasRole(roles[0], admin)

    # ---Revert path testing---

    # Only the role of default Admin has call permissions.
    with brownie .reverts():
        uni_iotx.revokeRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_clear.revokeRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_staking.revokeRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
