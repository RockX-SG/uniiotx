import brownie


def test_snapshot(fn_isolation, contracts, admin, users):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    # Minting uniIOTXs
    uni_iotx.mint(users[0], 100, {'from': iotx_staking, 'allow_revert': True})

    # Snapshot1
    tx = uni_iotx.snapshot({'from': iotx_staking, 'allow_revert': True})
    assert "Snapshot" in tx.events
    assert uni_iotx.totalSupplyAt(1) == 100
    assert uni_iotx.balanceOfAt(users[0], 1) == 100

    # Snapshot2
    tx = uni_iotx.snapshot({'from': iotx_staking, 'allow_revert': True})
    assert "Snapshot" in tx.events
    assert uni_iotx.totalSupplyAt(2) == 100
    assert uni_iotx.balanceOfAt(users[0], 2) == 100

    # ---Revert path testing---

    # Only the Minter role has call permission.
    with brownie .reverts():
        uni_iotx.snapshot({'from': admin, 'allow_revert': True})
