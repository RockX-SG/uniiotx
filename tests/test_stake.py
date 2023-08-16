import brownie


def test_stake(fn_isolation, contracts, stake_amounts, oracles, users, delegates, admin, deadline):
    system_staking, iotx_staking = contracts[0], contracts[3]

    # ---Happy path testing---

    # Depositing IOTXs
    amounts = iotx_staking.getStakeAmounts()
    amount_top = amounts[len(amounts) - 1]
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amount_top, 'allow_revert': True})

    # Mock reward
    mock_reward = amounts[1]
    delegates[0].transfer(iotx_staking, mock_reward)
    iotx_staking.updateReward({'from': oracles[0], 'allow_revert': True})
    user_reward = iotx_staking.getUserReward()
    manager_reward = iotx_staking.getManagerReward()

    # Normal staking
    tx = iotx_staking.stake({'from': oracles[0], 'allow_revert': True})
    token_id_slice = [2, 3, 4, 5, 6, 7, 8, 9, 10]
    assert "Staked" in tx.events
    assert system_staking.balance() == amount_top + user_reward
    assert iotx_staking.getTotalStaked() == amount_top + user_reward
    assert iotx_staking.getAccountedBalance() == manager_reward
    assert iotx_staking.getTotalPending() == 0
    assert iotx_staking.getTokenQueueLength(0) == 9
    assert iotx_staking.getStakedTokenCount(0) == 9
    assert iotx_staking.getStakedTokenIds(0) == token_id_slice
    assert iotx_staking.getStakedTokenIdSlice(0, 0, 9) == token_id_slice
    for i in range(0, 9):
        assert iotx_staking.getTokenId(0, i) == token_id_slice[i]

    # Merging operation
    iotx_staking.withdrawManagerFee(manager_reward, admin, {'from': admin, 'allow_revert': True})
    tx = iotx_staking.stake({'from': oracles[0], 'allow_revert': True})
    assert "Staked" in tx.events
    assert "Merged" in tx.events
    assert system_staking.balance() == amount_top + mock_reward
    assert iotx_staking.getTotalStaked() == amount_top + mock_reward
    assert iotx_staking.getAccountedBalance() == 0
    assert iotx_staking.getTotalPending() == 0
    assert iotx_staking.getTokenQueueLength(0) == 0
    assert iotx_staking.getStakedTokenCount(0) == 0
    assert iotx_staking.getTokenQueueLength(1) == 1
    assert iotx_staking.getStakedTokenCount(1) == 1
    assert iotx_staking.getStakedTokenIds(1) == [2]
    assert iotx_staking.getStakedTokenIdSlice(1, 0, 1) == [2]

    # ---Revert path testing---

    # When the contract is on pause, the 'stake' function will not operate.
    iotx_staking.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_staking.stake({'from': oracles[0], 'allow_revert': True})
    iotx_staking.unpause({'from': admin})

    # Only the role of oracle has call permissions.
    with brownie .reverts():
        iotx_staking.stake({'from': admin, 'allow_revert': True})
