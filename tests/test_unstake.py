import brownie


def test_unstake(fn_isolation, contracts, users, oracles, admin, deadline, uint256_max):
    system_staking, uni_iotx, iotx_clear, iotx_staking = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # Depositing IOTXs
    amounts = iotx_staking.getStakeAmounts()
    amount_top = amounts[len(amounts) - 1]
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amount_top, 'allow_revert': True})

    # Redeeming IOTXs
    uni_iotx.approve(iotx_staking, amount_top, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amount_top, deadline, {'from': users[0], 'allow_revert': True})

    # Unstaking the unlocked token
    token_id = iotx_staking.getTokenId(2, 0)
    tx = iotx_clear.unstake([token_id], {'from': oracles[0], 'allow_revert': True})
    bucket = system_staking.bucketOf(token_id)
    assert "Unstaked" in tx.events
    assert bucket[3] < uint256_max

    # ---Revert path testing---

    # When the contract is on pause, the 'unstake' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.unstake([token_id], {'from': oracles[0], 'allow_revert': True})
    iotx_clear.unpause({'from': admin})

    # Only the role of stake has call permission.
    with brownie .reverts():
        iotx_clear.unstake([token_id], {'from': admin, 'allow_revert': True})

    # Only debt tokens are accepted.
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amounts[0], 'allow_revert': True})
    with brownie .reverts("USR007"):
        iotx_clear.unstake([iotx_staking.getTokenId(0, 0)], {'from': oracles[0], 'allow_revert': True})
