from configs import *
from contracts import *

def test_init_status(delegates, start_amount, common_ratio, sequence_length, stake_amounts, stake_duration, manager_fee_shares, contracts):
    system_staking, uni_iotx, iotx_clear, iotx_stake = contracts[0], contracts[1], contracts[2], contracts[3]
    assert system_staking.UNSTAKE_FREEZE_BLOCKS() == 1
    assert uni_iotx.symbol() == "uniIOTX"
    assert iotx_clear.MULTIPLIER() == 1000000000000000000
    assert iotx_stake.defaultExchangeRatio() == 1
    assert iotx_stake.MULTIPLIER() == 1000000000000000000
    assert iotx_stake.globalDelegate() == delegates[0].address
    assert iotx_stake.startAmount() == start_amount
    assert iotx_stake.commonRatio() == common_ratio
    assert iotx_stake.sequenceLength() == sequence_length
    assert iotx_stake.redeemAmountBase() == stake_amounts[2]
    assert iotx_stake.stakeDuration() == stake_duration
    assert iotx_stake.managerFeeShares() == manager_fee_shares