import brownie


def test_mint(fn_isolation, contracts, users, zero_address):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 100
    tx = uni_iotx.mint(iotx_staking, amt, {'from': iotx_staking, 'allow_revert': True})
    assert 'Transfer' in tx.events
    assert uni_iotx.totalSupply() == amt
    assert uni_iotx.balanceOf(iotx_staking) == amt

    # ---Revert path testing---

    # Minting to a zero address is not allowed.
    with brownie .reverts("ERC20: mint to the zero address"):
        uni_iotx.mint(zero_address, amt, {'from': iotx_staking, 'allow_revert': True})

    # Only the Minter role has the authority to make calls.
    with brownie .reverts():
        uni_iotx.mint(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
