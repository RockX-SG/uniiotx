import brownie


def test_claim(fn_isolation, contracts, users, oracles, admin, delegates, deadline):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # After delegate distributes rewards, users' can claim their reward portion.
    # Users can claim any available reward amount at various times.
    # Once users have claimed their rewards, the individual reward record
    # and the global 'accountedBalance' should be adjusted accordingly.

    # After repaying their debt, users' can claim their principal.
    # Users can claim any available principal amount at various times.
    # Once users have claimed their principals, the individual principal record
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

    # Pay debt
    token_ids = iotx_staking.getRedeemedTokenIds()
    iotx_clear.unstake(token_ids, {'from': oracles[0], 'allow_revert': True})
    iotx_clear.payDebts(token_ids, {'from': oracles[0], 'allow_revert': True})

    # Claim partial rewards
    balance0 = users[1].balance()
    tx = iotx_clear.claim(mock_reward/2, users[1], {'from': users[0], 'allow_revert': True})
    assert "Claimed" in tx.events
    assert users[1].balance() == balance0 + mock_reward/2
    assert iotx_clear.balance() == amt + mock_reward/2
    assert iotx_clear.getAccountedBalance() == amt + mock_reward/2
    assert iotx_clear.getBalance(users[0]) == mock_reward/2 + amt
    assert iotx_clear.getReward(users[0]) == mock_reward/2
    assert iotx_clear.getPrincipal(users[0]) == amt
    assert iotx_clear.getUserInfo(users[0])[2] == mock_reward/2
    assert iotx_clear.getUserInfo(users[0])[1] == amt

    # Claim the rest of rewards and partial principal
    tx = iotx_clear.claim(mock_reward/2 + amt/2, users[1], {'from': users[0], 'allow_revert': True})
    assert "Claimed" in tx.events
    assert users[1].balance() == balance0 + mock_reward + amt/2
    assert iotx_clear.balance() == amt/2
    assert iotx_clear.getAccountedBalance() == amt/2
    assert iotx_clear.getBalance(users[0]) == amt/2
    assert iotx_clear.getReward(users[0]) == 0
    assert iotx_clear.getPrincipal(users[0]) == amt/2
    assert iotx_clear.getUserInfo(users[0])[2] == 0
    assert iotx_clear.getUserInfo(users[0])[1] == amt/2

    # Claim the rest of principal
    tx = iotx_clear.claim(amt/2, users[1], {'from': users[0], 'allow_revert': True})
    assert "Claimed" in tx.events
    assert users[1].balance() == balance0 + mock_reward + amt
    assert iotx_clear.balance() == 0
    assert iotx_clear.getAccountedBalance() == 0
    assert iotx_clear.getBalance(users[0]) == 0
    assert iotx_clear.getReward(users[0]) == 0
    assert iotx_clear.getPrincipal(users[0]) == 0
    assert iotx_clear.getUserInfo(users[0])[2] == 0
    assert iotx_clear.getUserInfo(users[0])[1] == 0

    # Submit a deposit request.
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})

    # Submit a redeem request for the creation of a debt item.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})

    # Mock reward.
    mock_reward = 100
    delegates[0].transfer(iotx_clear, mock_reward)
    # Pay debt
    token_ids = iotx_staking.getRedeemedTokenIdSlice(1, 2)
    iotx_clear.unstake(token_ids, {'from': oracles[0], 'allow_revert': True})
    iotx_clear.payDebts(token_ids, {'from': oracles[0], 'allow_revert': True})

    # Claim rewards and principal in one lump sum
    tx = iotx_clear.claim(amt + mock_reward, users[1], {'from': users[0], 'allow_revert': True})
    assert "Claimed" in tx.events
    assert users[1].balance() == balance0 + 2*mock_reward + 2*amt
    assert iotx_clear.balance() == 0
    assert iotx_clear.getAccountedBalance() == 0
    assert iotx_clear.getBalance(users[0]) == 0
    assert iotx_clear.getReward(users[0]) == 0
    assert iotx_clear.getPrincipal(users[0]) == 0
    assert iotx_clear.getUserInfo(users[0])[2] == 0
    assert iotx_clear.getUserInfo(users[0])[1] == 0

    # ---Revert path testing---

    # When the contract is on pause, the 'claim' function will not operate.
    iotx_clear.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_clear.claim(amt/2, users[0], {'from': users[0]})
    iotx_clear.unpause({'from': admin})

    # Users should possess sufficient asset
    with brownie .reverts("USR009"):
        iotx_clear.claim(amt/2, users[0], {'from': users[0]})

    # Todo: Handle nonReentrant
