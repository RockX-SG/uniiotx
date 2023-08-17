import brownie


def test_transfer(fn_isolation, contracts, users, zero_address, uint256_max):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    # With finite approval
    amt = 100
    uni_iotx.mint(users[0], amt, {'from': iotx_staking, 'allow_revert': True})
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    tx = uni_iotx.transferFrom(users[0], users[1], amt, {'from': iotx_staking, 'allow_revert': True})
    assert "Approval" in tx.events
    assert uni_iotx.allowance(users[0], iotx_staking) == 0
    assert "Transfer" in tx.events
    assert uni_iotx.balanceOf(users[1]) == amt

    # With infinite approval
    uni_iotx.mint(users[0], amt, {'from': iotx_staking, 'allow_revert': True})
    uni_iotx.approve(iotx_staking, uint256_max, {'from': users[0], 'allow_revert': True})
    tx = uni_iotx.transferFrom(users[0], users[1], amt, {'from': iotx_staking, 'allow_revert': True})
    assert "Approval" not in tx.events
    assert uni_iotx.allowance(users[0], iotx_staking) == uint256_max
    assert "Transfer" in tx.events
    assert uni_iotx.balanceOf(users[1]) == amt*2

    # ---Revert path testing---

    # Not allowed to approve from the zero address
    uni_iotx.approve(iotx_staking, 0, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("ERC20: insufficient allowance"):
        uni_iotx.transferFrom(users[0], users[1], amt, {'from': iotx_staking, 'allow_revert': True})

    # Not allowed to approve from the zero address
    with brownie .reverts("ERC20: approve from the zero address"):
        uni_iotx.transferFrom(zero_address, users[1], 0, {'from': iotx_staking, 'allow_revert': True})

    # Not allowed to approve to the zero address
    with brownie .reverts("ERC20: approve to the zero address"):
        uni_iotx.transferFrom(users[0],  users[1], 0, {'from': zero_address, 'allow_revert': True})

    # Not allowed to transfer to the zero address
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("ERC20: transfer to the zero address"):
        uni_iotx.transferFrom(users[0], zero_address, amt, {'from': iotx_staking, 'allow_revert': True})

    # Transfer amount shouldn't exceed balance
    with brownie .reverts("ERC20: transfer amount exceeds balance"):
        uni_iotx.transferFrom(users[0], users[1], amt, {'from': iotx_staking, 'allow_revert': True})
