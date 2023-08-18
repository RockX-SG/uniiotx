import brownie


def test_burn(fn_isolation, contracts, users, roles, admin, zero_address):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 100
    uni_iotx.mint(users[0], amt, {'from': iotx_staking, 'allow_revert': True})
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})

    tx = uni_iotx.burnFrom(users[0], amt/2, {'from': iotx_staking, 'allow_revert': True})
    assert 'Approval' in tx.events
    assert 'Transfer' in tx.events
    assert uni_iotx.totalSupply() == amt/2
    assert uni_iotx.balanceOf(users[0]) == amt/2
    assert uni_iotx.allowance(users[0], iotx_staking) == amt/2

    # ---Revert path testing---

    # Insufficient allowance
    with brownie .reverts("ERC20: insufficient allowance"):
        uni_iotx.burnFrom(users[0], amt, {'from': iotx_staking, 'allow_revert': True})

    # Not allowed to approve from the zero address
    with brownie .reverts("ERC20: approve from the zero address"):
        uni_iotx.burnFrom(zero_address, 0, {'from': iotx_staking, 'allow_revert': True})

    # Not allowed to approve to the zero address
    uni_iotx.grantRole(roles[1], zero_address, {'from': admin, 'allow_revert': True})
    with brownie .reverts("ERC20: approve to the zero address"):
        uni_iotx.burnFrom(users[0], 0, {'from': zero_address, 'allow_revert': True})

    # The account should maintain an adequate balance.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("ERC20: burn amount exceeds balance"):
        uni_iotx.burnFrom(users[0], amt, {'from': iotx_staking, 'allow_revert': True})

    # Only the Minter role has the authority to make calls.
    with brownie .reverts():
        uni_iotx.burnFrom(users[0], amt, {'from': admin, 'allow_revert': True})
