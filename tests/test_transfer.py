import brownie


def test_pause(fn_isolation, contracts, users, zero_address):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 100
    uni_iotx.mint(users[0], amt, {'from': iotx_staking, 'allow_revert': True})
    tx = uni_iotx.transfer(users[1], amt, {'from': users[0], 'allow_revert': True})
    assert "Transfer" in tx.events
    assert uni_iotx.balanceOf(users[1]) == amt

    # ---Revert path testing---

    # Not allowed to transfer from the zero address
    with brownie .reverts("ERC20: transfer from the zero address"):
        uni_iotx.transfer(users[1], amt, {'from': zero_address, 'allow_revert': True})

    # Not allowed to transfer to the zero address
    with brownie .reverts("ERC20: transfer to the zero address"):
        uni_iotx.transfer(zero_address, amt, {'from': users[0], 'allow_revert': True})

    # Transfer amount shouldn't exceed balance
    with brownie .reverts("ERC20: transfer amount exceeds balance"):
        uni_iotx.transfer(users[1], amt, {'from': users[0], 'allow_revert': True})
