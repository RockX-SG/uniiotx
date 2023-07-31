from configs import *
from contracts import *

def test_init_status(roles, owner, admin, delegates, oracle, start_amount, common_ratio, sequence_length, stake_amounts, stake_duration, manager_fee_shares, contracts):
    system_staking, uni_iotx, iotx_clear, iotx_stake = contracts[0], contracts[1], contracts[2], contracts[3]

    # Check permissions
    assert system_staking.owner() == owner.address

    assert uni_iotx.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert uni_iotx.hasRole(roles[0], iotx_stake.address)  # ROLE_PAUSE
    assert uni_iotx.hasRole(roles[1], iotx_stake.address)  # ROLE_MING

    assert iotx_clear.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert iotx_clear.hasRole(roles[2], iotx_stake.address)  # ROLE_STAKE
    assert iotx_clear.hasRole(roles[4], oracle.address)  # ROLE_ORACLE

    assert iotx_stake.hasRole(roles[5], admin.address)  # ROLE_DEFAULT_ADMIN
    assert iotx_stake.hasRole(roles[3], admin.address)  # ROLE_FEE_MANAGER
    assert iotx_stake.hasRole(roles[0], admin.address)  # ROLE_PAUSE
    assert iotx_stake.hasRole(roles[4], oracle.address)  # ROLE_ORACLE

    # Check initial status variables
    assert system_staking.UNSTAKE_FREEZE_BLOCKS() == 1

    assert uni_iotx.symbol() == "uniIOTX"

    assert iotx_clear.MULTIPLIER() == 1000000000000000000

    assert iotx_stake.DEFAULT_EXCHANGE_RATIO() == 1
    assert iotx_stake.MULTIPLIER() == 1000000000000000000
    assert iotx_stake.globalDelegate() == delegates[0].address
    assert iotx_stake.startAmount() == start_amount
    assert iotx_stake.commonRatio() == common_ratio
    assert iotx_stake.sequenceLength() == sequence_length
    assert iotx_stake.redeemAmountBase() == stake_amounts[2]
    assert iotx_stake.stakeDuration() == stake_duration
    assert iotx_stake.managerFeeShares() == manager_fee_shares
