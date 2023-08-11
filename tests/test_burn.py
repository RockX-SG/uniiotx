import brownie


def test_burn(fn_isolation, contracts, users):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 100
    tx = uni_iotx.mint(iotx_staking, amt, {'from': iotx_staking, 'allow_revert': True})
    assert 'Transfer' in tx.events
    assert uni_iotx.totalSupply() == amt
    assert uni_iotx.balanceOf(iotx_staking) == amt

    tx = uni_iotx.burn(amt, {'from': iotx_staking, 'allow_revert': True})
    assert 'Transfer' in tx.events
    assert uni_iotx.totalSupply() == 0
    assert uni_iotx.balanceOf(iotx_staking) == 0

    # ---Revert path testing---

    # The account should maintain an adequate balance.
    with brownie .reverts("ERC20: burn amount exceeds balance"):
        uni_iotx.burn(amt, {'from': iotx_staking, 'allow_revert': True})

    # Only the Minter role has the authority to make calls.
    with brownie .reverts():
        uni_iotx.burn(amt, {'from': users[0], 'allow_revert': True})
