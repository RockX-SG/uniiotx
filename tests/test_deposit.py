import brownie


def test_deposit(fn_isolation, w3, contracts, stake_amounts, users, admin):
    system_staking, uni_iotx, iotx_staking = contracts[0], contracts[1], contracts[3]

    # ---Happy path testing---

    # Any deposit amount that is smaller than the 'startAmount' should be acceptable.
    deadline = w3.eth.get_block('latest').timestamp+60
    start_amt = iotx_staking.startAmount()
    small_amt = start_amt / 10
    tx = iotx_staking.deposit(deadline, {'from': users[0], 'value': small_amt, 'allow_revert': True})
    assert iotx_staking.accountedBalance() == small_amt
    assert iotx_staking.balance() == small_amt
    assert iotx_staking.totalPending() == small_amt
    assert iotx_staking.totalStaked() == 0
    assert uni_iotx.totalSupply() == small_amt
    assert uni_iotx.balanceOf(users[0]) == small_amt
    assert "Minted" in tx.events
    assert "Staked" not in tx.events
    assert system_staking.balanceOf(iotx_staking) == 0

    # Any deposit request should automatically trigger SystemStaking
    # if it causes the 'totalPending' exceed the 'startAmount'
    extra_amt = start_amt - small_amt
    tx = iotx_staking.deposit(deadline, {'from': users[0], 'value': extra_amt, 'allow_revert': True})
    assert iotx_staking.accountedBalance() == 0
    assert iotx_staking.balance() == 0
    assert iotx_staking.totalPending() == 0
    assert iotx_staking.totalStaked() == start_amt
    assert iotx_staking.getStakedTokenCount(0) == 1
    assert uni_iotx.totalSupply() == start_amt
    assert uni_iotx.balanceOf(users[0]) == start_amt
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert len(tx.events["Staked"]) == 2
    token_id = iotx_staking.getTokenId(0, 0)
    amt, dur, _, _, delegate = system_staking.bucketOf(token_id)
    assert (amt, dur, delegate) == (start_amt, iotx_staking.stakeDuration(), iotx_staking.globalDelegate())
    assert system_staking.balanceOf(iotx_staking) == 1
    assert system_staking.ownerOf(token_id) == iotx_staking

    # The merge operation at low staking level should be triggered
    # if the amount of staked tokens reaches the 'commonRatio'
    for i in range(0, iotx_staking.commonRatio() - 2):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': start_amt, 'allow_revert': True})
    tx = iotx_staking.deposit(deadline, {'from': users[0], 'value': start_amt, 'allow_revert': True})
    assert iotx_staking.accountedBalance() == 0
    assert iotx_staking.balance() == 0
    assert iotx_staking.totalPending() == 0
    assert iotx_staking.totalStaked() == stake_amounts[1]
    assert iotx_staking.getStakedTokenCount(0) == 0
    assert iotx_staking.getStakedTokenCount(1) == 1
    assert uni_iotx.totalSupply() == stake_amounts[1]
    assert uni_iotx.balanceOf(users[0]) == stake_amounts[1]
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert len(tx.events["Staked"]) == 2
    assert token_id == iotx_staking.getTokenId(1, 0)
    amt, dur, _, _, delegate = system_staking.bucketOf(token_id)
    assert (amt, dur, delegate) == (stake_amounts[1], iotx_staking.stakeDuration(), iotx_staking.globalDelegate())
    assert system_staking.balanceOf(iotx_staking) == 1
    assert system_staking.ownerOf(token_id) == iotx_staking

    for i in range(0, iotx_staking.commonRatio() - 2):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[1], 'allow_revert': True})
    tx = iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[1], 'allow_revert': True})
    assert iotx_staking.accountedBalance() == 0
    assert iotx_staking.balance() == 0
    assert iotx_staking.totalPending() == 0
    assert iotx_staking.totalStaked() == stake_amounts[2]
    assert iotx_staking.getStakedTokenCount(1) == 0
    assert iotx_staking.getStakedTokenCount(2) == 1
    assert uni_iotx.totalSupply() == stake_amounts[2]
    assert uni_iotx.balanceOf(users[0]) == stake_amounts[2]
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert len(tx.events["Staked"]) == 2
    assert token_id == iotx_staking.getTokenId(2, 0)
    amt, dur, _, _, delegate = system_staking.bucketOf(token_id)
    assert (amt, dur, delegate) == (stake_amounts[2], iotx_staking.stakeDuration(), iotx_staking.globalDelegate())
    assert system_staking.balanceOf(iotx_staking) == 1
    assert system_staking.ownerOf(token_id) == iotx_staking

    # There should be no merge operation at the top staking level.
    for i in range(0, iotx_staking.commonRatio() - 2):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[2], 'allow_revert': True})
    tx = iotx_staking.deposit(deadline, {'from': users[0], 'value': stake_amounts[2], 'allow_revert': True})
    assert iotx_staking.accountedBalance() == 0
    assert iotx_staking.balance() == 0
    assert iotx_staking.totalPending() == 0
    assert iotx_staking.totalStaked() == stake_amounts[2] * iotx_staking.commonRatio()
    assert iotx_staking.getStakedTokenCount(2) == 10
    user_balance = stake_amounts[2] * iotx_staking.commonRatio()
    assert uni_iotx.totalSupply() == user_balance
    assert uni_iotx.balanceOf(users[0]) == user_balance
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert len(tx.events["Staked"]) == 2
    assert system_staking.balanceOf(iotx_staking) == 10
    for i in range(0, 10):
        token_id = iotx_staking.getTokenId(2, i)
        amt, dur, _, _, delegate = system_staking.bucketOf(token_id)
        assert (amt, dur, delegate) == (stake_amounts[2], iotx_staking.stakeDuration(), iotx_staking.globalDelegate())
        assert system_staking.ownerOf(token_id) == iotx_staking

    # ---Revert path testing---

    # When the contract is on pause, the 'deposit' function will not operate.
    iotx_staking.pause({'from': admin})
    with brownie .reverts("Pausable: paused"):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': start_amt, 'allow_revert': True})
    iotx_staking.unpause({'from': admin})
    iotx_staking.deposit(deadline, {'from': users[0], 'value': start_amt, 'allow_revert': True})

    # The transaction of the deposit request should arrive within the deadline time.
    past_deadline = "1690514039"
    with brownie .reverts("USR001"):
        iotx_staking.deposit(past_deadline, {'from': users[0], 'value': start_amt, 'allow_revert': True})

    # Deposits of zero value are not permitted.
    with brownie .reverts("USR002"):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': 0, 'allow_revert': True})

    # Todo: Handle nonReentrant



