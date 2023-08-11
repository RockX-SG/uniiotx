import brownie


def test_claimRewards(fn_isolation, contracts, users, admin, delegates, deadline):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # After delegate distributes rewards, users' can claim their reward portion.
    # Users can claim any available reward amount at various times.
    # Once users have claimed their rewards, the individual reward record
    # and the global 'accountedBalance' should be adjusted accordingly.

    # Submit a deposit request.
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})

    # Submit a redeem request for the creation of a debt item.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})

    # Mock reward.
    mock_reward = 100
    delegates[0].transfer(iotx_clear, mock_reward)

    # Claim user rewards.
    balance0 = users[1].balance()
    for i in range(0, 2):
        tx = iotx_clear.claimRewards(mock_reward/2, users[1], {'from': users[0], 'allow_revert': True})
        assert "RewardClaimed" in tx.events
    assert users[1].balance() == balance0 + mock_reward
    assert iotx_clear.balance() == 0
    assert iotx_clear.getAccountedBalance() == 0
    assert iotx_clear.getReward(users[0]) == 0
    assert iotx_clear.getUserInfo(users[0])[2] == 0

    # ---Revert path testing---

    # When the contract is on pause, the 'claimRewards' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.claimRewards(amt/2, users[0], {'from': users[0]})
    iotx_clear.unpause({'from': admin})

    # Users should possess sufficient rewards
    with brownie .reverts("USR005"):
        iotx_clear.claimRewards(amt/2, users[0], {'from': users[0]})

    # Todo: Handle nonReentrant
