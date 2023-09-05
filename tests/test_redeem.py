import brownie


def test_redeem(fn_isolation, contracts, users, delegates, oracles, deadline, uint256_max):
    system_staking, uni_iotx, iotx_clear, iotx_staking = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    amt = iotx_staking.getRedeemAmountBase()
    for i in range(0, 2):
        iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})
    amt_reward = 100
    delegates[0].transfer(iotx_staking, amt_reward)
    assert iotx_staking.getPendingReward() == amt_reward

    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    tx = iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})
    token_id = iotx_staking.getTokenId(2, 0)
    unlocked_amt, _, unlocked_at, _, _ = system_staking.bucketOf(token_id)
    debt = iotx_clear.getUnpaidDebtItem(1)
    user_info = iotx_clear.getUserInfo(users[0])
    assert len(tx.events["Transfer"]) == 2
    assert len(tx.events["Unlocked"]) == 1
    assert len(tx.events["Redeemed"]) == 1
    assert len(tx.events["DebtQueued"]) == 1
    assert unlocked_amt == amt
    assert unlocked_at != uint256_max
    assert iotx_staking.getRedeemedTokenCount() == 1
    assert iotx_staking.getTotalStaked() == amt
    assert iotx_staking.getPendingReward() == 0
    assert uni_iotx.totalSupply() == 10045
    assert uni_iotx.balanceOf(users[0]) == 10045
    assert system_staking.balanceOf(iotx_staking) == 1
    assert system_staking.ownerOf(token_id) == iotx_clear
    assert system_staking.balanceOf(iotx_clear) == 1
    assert iotx_clear.getTotalDebts() == amt
    assert debt[0] == users[0]
    assert debt[1] == amt
    assert user_info[0] == amt

    # ---Revert path testing---

    # The transaction of the redeem request should arrive within the deadline time.
    past_deadline = "1690514039"
    with brownie .reverts("USR001"):
        iotx_staking.redeem(amt, past_deadline, {'from': users[0], 'allow_revert': True})

    # Users can can only redeem amounts that are multiples of the redeemAmountBase.
    with brownie .reverts("USR003"):
        iotx_staking.redeem(amt - 10, deadline, {'from': users[0], 'allow_revert': True})
    with brownie .reverts("USR003"):
        iotx_staking.redeem(amt + 10, deadline, {'from': users[0], 'allow_revert': True})

    # Todo: Handle nonReentrant
    # Todo: Handle "ERC20: insufficient allowance"
