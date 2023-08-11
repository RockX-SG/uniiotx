import brownie


def test_getReward(fn_isolation, contracts, users, delegates, deadline):
    uni_iotx, iotx_clear, iotx_staking = contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # Submit a deposit request
    amt = iotx_staking.getRedeemAmountBase()
    iotx_staking.deposit(deadline, {'from': users[0], 'value': amt, 'allow_revert': True})

    # Submit a redeem request for the creation of a debt item.
    uni_iotx.approve(iotx_staking, amt, {'from': users[0], 'allow_revert': True})
    iotx_staking.redeem(amt, deadline, {'from': users[0], 'allow_revert': True})

    # Rewards will not be credited if the account has no outstanding debt.
    assert iotx_clear.getReward(users[1]) == 0

    # Rewards will not be credited unless the delegates distribute them.
    assert iotx_clear.getReward(users[0]) == 0
    mock_reward = 100
    delegates[0].transfer(iotx_clear, mock_reward)
    assert iotx_clear.getReward(users[0]) == mock_reward

    # Todo: add more test cases.
