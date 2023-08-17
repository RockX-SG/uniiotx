import brownie


def test_increaseAllowance(fn_isolation, contracts, users, zero_address, uint256_max):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 100
    tx = uni_iotx.increaseAllowance(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    assert "Approval" in tx.events
    assert uni_iotx.allowance(users[0], iotx_staking) == amt

    # ---Revert path testing---

    # Not allowed to approve from the zero address
    with brownie .reverts("ERC20: approve from the zero address"):
        uni_iotx.approve(iotx_staking, amt, {'from': zero_address, 'allow_revert': True})

    # Not allowed to approve to the zero address
    with brownie .reverts("ERC20: approve to the zero address"):
        uni_iotx.approve(zero_address, amt, {'from': users[0], 'allow_revert': True})

    # Not allow integer overflow
    with brownie .reverts("Integer overflow"):
        uni_iotx.increaseAllowance(iotx_staking, uint256_max, {'from': users[0], 'allow_revert': True})
