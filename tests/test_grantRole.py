import brownie


def test_grantRole(fn_isolation, contracts, users, roles, admin, oracles):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # First time grant
    tx = uni_iotx.grantRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleGranted" in tx.events
    assert uni_iotx.hasRole(roles[0], admin)
    assert uni_iotx.getRoleAdmin(roles[0]) == "0x00"

    tx = iotx_clear.grantRole(roles[0], oracles[0], {'from': admin, 'allow_revert': True})
    assert "RoleGranted" in tx.events
    assert iotx_clear.hasRole(roles[0], oracles[0])
    assert iotx_clear.getRoleAdmin(roles[0]) == "0x00"

    tx = iotx_staking.grantRole(roles[0], oracles[0], {'from': admin, 'allow_revert': True})
    assert "RoleGranted" in tx.events
    assert iotx_staking.hasRole(roles[0], oracles[0])
    assert iotx_staking.getRoleAdmin(roles[0]) == "0x00"

    # No repeated grant
    tx = uni_iotx.grantRole(roles[0], iotx_staking, {'from': admin, 'allow_revert': True})
    assert "RoleGranted" not in tx.events
    assert uni_iotx.hasRole(roles[0], iotx_staking)
    assert uni_iotx.getRoleAdmin(roles[0]) == "0x00"

    tx = iotx_clear.grantRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleGranted" not in tx.events
    assert iotx_clear.hasRole(roles[0], admin)
    assert iotx_clear.getRoleAdmin(roles[0]) == "0x00"

    tx = iotx_staking.grantRole(roles[0], admin, {'from': admin, 'allow_revert': True})
    assert "RoleGranted" not in tx.events
    assert iotx_staking.hasRole(roles[0], admin)
    assert iotx_staking.getRoleAdmin(roles[0]) == "0x00"

    # ---Revert path testing---

    # Only the role of default Admin has call permissions.
    with brownie .reverts():
        uni_iotx.grantRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_clear.grantRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
    with brownie .reverts():
        iotx_staking.grantRole(roles[0], iotx_staking, {'from': users[0], 'allow_revert': True})
