import brownie


def test_init_status(fn_isolation, roles, owner, admin, delegates, oracles, start_amount, common_ratio, sequence_length, stake_amounts, stake_duration, manager_fee_shares, contracts):
    system_staking, uni_iotx, iotx_clear, iotx_staking = contracts[0], contracts[1], contracts[2], contracts[3]

    # ---Happy path testing---

    # Check permissions
    assert system_staking.owner() == owner.address

    assert uni_iotx.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert uni_iotx.hasRole(roles[0], iotx_staking.address)  # ROLE_PAUSE
    assert uni_iotx.hasRole(roles[1], iotx_staking.address)  # ROLE_MING

    assert iotx_clear.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert iotx_clear.hasRole(roles[2], iotx_staking.address)  # ROLE_STAKE
    assert iotx_clear.hasRole(roles[4], oracles[0].address)  # ROLE_ORACLE

    assert iotx_staking.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert iotx_staking.hasRole(roles[3], admin.address)  # ROLE_FEE_MANAGER
    assert iotx_staking.hasRole(roles[0], admin.address)  # ROLE_PAUSE
    assert iotx_staking.hasRole(roles[4], oracles[0].address)  # ROLE_ORACLE

    # Check initial status variables
    assert system_staking.UNSTAKE_FREEZE_BLOCKS() == 1

    assert uni_iotx.name() == "Universal IOTX"
    assert uni_iotx.symbol() == "uniIOTX"
    assert uni_iotx.decimals() == 18

    assert iotx_clear.getDebtAmountBase() == iotx_staking.getRedeemAmountBase()

    assert iotx_staking.getGlobalDelegate() == delegates[0].address
    assert iotx_staking.startAmount() == start_amount
    assert iotx_staking.commonRatio() == common_ratio
    assert iotx_staking.sequenceLength() == sequence_length
    assert iotx_staking.getRedeemAmountBase() == stake_amounts[2]
    assert iotx_staking.getStakeDuration() == stake_duration
    assert iotx_staking.getManagerFeeShares() == manager_fee_shares

    # ---Revert path testing---

    # Reinitialization is not allowed
    with brownie .reverts("Initializable: contract is already initialized"):
        uni_iotx.initialize(iotx_staking, {'from': admin})
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_staking.initialize(
            system_staking,
            uni_iotx,
            iotx_clear,
            oracles[0],
            start_amount,
            common_ratio,
            sequence_length,
            stake_duration,
            {'from': admin}
        )
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_clear.initialize(system_staking, iotx_staking, oracles[0], {'from': admin})
