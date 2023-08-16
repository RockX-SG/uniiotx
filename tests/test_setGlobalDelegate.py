import brownie


def test_setGlobalDelegate(fn_isolation, contracts, delegates, oracles, admin):
    iotx_staking = contracts[3]

    # ---Happy path testing---

    tx = iotx_staking.setGlobalDelegate(delegates[1], {'from': oracles[0]})
    assert "GlobalDelegateSet" in tx.events
    assert iotx_staking.getGlobalDelegate() == delegates[1]

    # ---Revert path testing---

    # When the contract is on pause, the 'setGlobalDelegate' function will not operate.
    iotx_staking.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_staking.setGlobalDelegate(delegates[1], {'from': oracles[0]})
    iotx_staking.unpause({'from': admin})

    # Only the Oracle role have call permissions.
    with brownie .reverts():
        iotx_staking.setGlobalDelegate(delegates[1], {'from': admin})
