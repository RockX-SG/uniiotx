from brownie import UniIOTX, IOTXClear, IOTXStaking, Contract
from web3 import Web3

# The command to run this script: `brownie run scripts/deploy/mainnet_check_init_status.py  --network=iotex-mainnet`

def main():
    # Contracts deployed
    uni_iotx = Contract.from_abi("UniIOTX", "0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894", UniIOTX.abi)
    iotx_clear = Contract.from_abi("IOTXClear", "0x7AD800771743F4e29f55235A55895273035FB546", IOTXClear.abi)
    iotx_staking = Contract.from_abi("IOTXStaking", "0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f", IOTXStaking.abi)

    # Bucket info
    start_amount = 10000 * 1e18
    common_ratio = 10
    sequence_length = 3
    stake_duration = 1572480  # (91 * 24 * 60 * 60) / 5

    # Manager fee shares
    manager_fee_shares = 50

    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.testnet.iotex.io'))

    # Contract roles
    role_pauser = w3.keccak(text='ROLE_PAUSER')
    role_minter = w3.keccak(text='ROLE_MINTER')
    role_staker = w3.keccak(text='ROLE_STAKER')
    role_fee_manager = w3.keccak(text='ROLE_FEE_MANAGER')
    role_oracle = w3.keccak(text='ROLE_ORACLE')
    role_default_admin = w3.toBytes(hexstr="0x00")

    # Accounts
    admin = "0xbFdDf5e269C74157b157c7DaC5E416d44afB790d"
    oracle = "0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09"
    delegate = "0xFE82234dE6b7F4DAE188552dfAa64C1dF5674caA"

    # Check permissions

    assert uni_iotx.hasRole(role_default_admin, admin)  # ROLE_DEFAULT_ADMIN
    assert uni_iotx.hasRole(role_pauser, iotx_staking.address)  # ROLE_PAUSER
    assert uni_iotx.hasRole(role_minter, iotx_staking.address)  # ROLE_MINTER

    assert iotx_clear.hasRole(role_default_admin, admin)  # ROLE_DEFAULT_ADMIN
    assert iotx_clear.hasRole(role_staker, iotx_staking.address)  # ROLE_STAKER
    assert iotx_clear.hasRole(role_oracle, oracle)  # ROLE_ORACLE

    assert iotx_staking.hasRole(role_default_admin, admin)  # ROLE_DEFAULT_ADMIN
    assert iotx_staking.hasRole(role_fee_manager, admin)  # ROLE_FEE_MANAGER
    assert iotx_staking.hasRole(role_pauser, admin)  # ROLE_PAUSER
    assert iotx_staking.hasRole(role_oracle, oracle)  # ROLE_ORACLE

    # Check initial status variables

    assert uni_iotx.name() == "Universal IOTX"
    assert uni_iotx.symbol() == "uniIOTX"
    assert uni_iotx.decimals() == 18

    assert iotx_clear.getDebtAmountBase() == iotx_staking.getRedeemAmountBase()

    assert iotx_staking.getGlobalDelegate() == delegate
    assert iotx_staking.startAmount() == start_amount
    assert iotx_staking.commonRatio() == common_ratio
    assert iotx_staking.sequenceLength() == sequence_length
    assert iotx_staking.getRedeemAmountBase() == start_amount * common_ratio * common_ratio
    assert iotx_staking.getStakeDuration() == stake_duration
    assert iotx_staking.getManagerFeeShares() == manager_fee_shares

