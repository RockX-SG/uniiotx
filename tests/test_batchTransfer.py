import brownie


def test_batchTransfer(fn_isolation, contracts, users):
    uni_iotx, iotx_staking = contracts[1], contracts[3]

    # ---Happy path testing---

    amt = 200
    uni_iotx.mint(iotx_staking, amt, {'from': iotx_staking, 'allow_revert': True})
    tx = uni_iotx.batchTransfer([users[0], users[1]], [amt/2, amt/2], {'from': iotx_staking, 'allow_revert': True})
    assert 'Transfer' in tx.events
    assert uni_iotx.totalSupply() == amt
    assert uni_iotx.balanceOf(iotx_staking) == 0
    assert uni_iotx.balanceOf(users[0]) == amt/2
    assert uni_iotx.balanceOf(users[1]) == amt/2

    # ---Revert path testing---

    # At least one recipient is provided.
    with brownie .reverts("uniIOTX: least one recipient address"):
        uni_iotx.batchTransfer([], [], {'from': iotx_staking, 'allow_revert': True})

    # The number of recipients must equal the number of tokens
    with brownie .reverts("uniIOTX: number of recipient addresses does not match the number of tokens"):
        uni_iotx.batchTransfer([users[0], users[1]], [amt/2], {'from': iotx_staking, 'allow_revert': True})
    with brownie .reverts("uniIOTX: number of recipient addresses does not match the number of tokens"):
        uni_iotx.batchTransfer([users[0]], [amt/2, amt/2], {'from': iotx_staking, 'allow_revert': True})
