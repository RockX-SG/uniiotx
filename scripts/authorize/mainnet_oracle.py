from brownie import UniIOTX, IOTXClear, IOTXStaking, accounts
from web3 import Web3

# The command to run this script: `brownie run scripts/authorize/mainnet_oracle.py  --network=iotex-mainnet`

# Executing the following scripts will:
# 1. Assign the ROLE_ORACLE to the Company's Oracle account, concurrently renouncing this permission from
#   the Personal Oracle account.

# The Personal Oracle account were temporarily used for contract deployment and initialization.
# Subsequently, the Company Oracle account is utilized during the production phase.


def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.mainnet.iotex.io'))

    # Contract roles
    role_pauser = w3.keccak(text='ROLE_PAUSER')
    role_minter = w3.keccak(text='ROLE_MINTER')
    role_staker = w3.keccak(text='ROLE_STAKER')
    role_fee_manager = w3.keccak(text='ROLE_FEE_MANAGER')
    role_oracle = w3.keccak(text='ROLE_ORACLE')
    role_default_admin = w3.toBytes(hexstr="0x00")

    # Contracts, managed by the following personal Admin account
    iotx_clear = IOTXClear.at("0x7AD800771743F4e29f55235A55895273035FB546")
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    # Personal Admin account
    personal_admin = accounts.load("IoTeXAdmin")

    # Personal Oracle account
    personal_oracle = accounts.load("IoTeXOracle")

    # Fake company Oracle account
    company_oracle = "0x5e3cD45341893E8129941Ea3f782BE8D68CE308d"

    # -------------------- IOTXClear --------------------
    # Permission transferred:: ROLE_ORACLE (Personal Oracle --> Company Oracle)
    if not iotx_clear.hasRole(role_oracle, company_oracle):
        iotx_clear.grantRole(role_oracle, company_oracle, {'from': personal_admin})
    assert iotx_clear.hasRole(role_oracle, company_oracle)

    if iotx_clear.hasRole(role_oracle, personal_oracle):
        iotx_clear.renounceRole(role_oracle, personal_oracle, {'from': personal_oracle})
    assert not iotx_clear.hasRole(role_oracle, personal_oracle)

    # -------------------- IOTXStaking --------------------
    # Permission transferred:: ROLE_ORACLE (Personal Oracle --> Company Oracle)
    if not iotx_staking.hasRole(role_oracle, company_oracle):
        iotx_staking.grantRole(role_oracle, company_oracle, {'from': personal_admin})
    assert iotx_staking.hasRole(role_oracle, company_oracle)

    if iotx_staking.hasRole(role_oracle, personal_oracle):
        iotx_staking.renounceRole(role_oracle, personal_oracle, {'from': personal_oracle})
    assert not iotx_staking.hasRole(role_oracle, personal_oracle)

