from brownie import UniIOTX, IOTXClear, IOTXStaking, accounts
from web3 import Web3

# The command to run this script: `brownie run scripts/authorize/mainnet_default_admin.py  --network=iotex-mainnet`

# Executing the following scripts will:
# 1. Grant the Company's Admin account the roles of ROLE_PAUSER, ROLE_FEE_MANAGER, and DEFAULT_ADMIN_ROLE,
#   while simultaneously revoking these permissions from the Personal Admin account.
# 2. Assign the ROLE_ORACLE to the Company's Oracle account, concurrently renouncing this permission from
#   the Personal Oracle account.

# The Personal Admin and Oracle accounts are temporarily used for contract deployment and initialization.
# Subsequently, the Company Admin and Oracle accounts are utilized during the production phase.


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
    uni_iotx = UniIOTX.at("0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894")
    iotx_clear = IOTXClear.at("0x7AD800771743F4e29f55235A55895273035FB546")
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    # Personal Admin account
    personal_admin = accounts.load("IoTeXAdmin")

    # Company Admin account, which will take over the aforementioned contracts
    company_admin = "0xE20aC9B2889EF2e4d9B6A4678e9be8d6048861B8"

    # -------------------- UniIOTX --------------------
    # Permission transferred: DEFAULT_ADMIN_ROLE (Personal Admin --> Company Admin)
    if not uni_iotx.hasRole(role_default_admin, company_admin):
        uni_iotx.grantRole(role_default_admin, company_admin, {'from': personal_admin})
    assert uni_iotx.hasRole(role_default_admin, company_admin)

    if uni_iotx.hasRole(role_default_admin, personal_admin):
        uni_iotx.renounceRole(role_default_admin, personal_admin, {'from': personal_admin})
    assert not uni_iotx.hasRole(role_default_admin, personal_admin)

    # -------------------- IOTXClear --------------------
    # Permission transferred:: ROLE_PAUSER (Personal Admin --> Company Admin)
    if not iotx_clear.hasRole(role_pauser, company_admin):
        iotx_clear.grantRole(role_pauser, company_admin, {'from': personal_admin})
    assert iotx_clear.hasRole(role_pauser, company_admin)

    if iotx_clear.hasRole(role_pauser, personal_admin):
        iotx_clear.renounceRole(role_pauser, personal_admin, {'from': personal_admin})
    assert not iotx_clear.hasRole(role_pauser, personal_admin)

    # Permission transferred:: DEFAULT_ADMIN_ROLE (Personal Admin --> Company Admin)
    if not iotx_clear.hasRole(role_default_admin, company_admin):
        iotx_clear.grantRole(role_default_admin, company_admin, {'from': personal_admin})
    assert iotx_clear.hasRole(role_default_admin, company_admin)

    if iotx_clear.hasRole(role_default_admin, personal_admin):
        iotx_clear.renounceRole(role_default_admin, personal_admin, {'from': personal_admin})
    assert not iotx_clear.hasRole(role_default_admin, personal_admin)

    # -------------------- IOTXStaking --------------------
    # Permission transferred:: ROLE_FEE_MANAGER (Personal Admin --> Company Admin)
    if not iotx_staking.hasRole(role_fee_manager, company_admin):
        iotx_staking.grantRole(role_fee_manager, company_admin, {'from': personal_admin})
    assert iotx_staking.hasRole(role_fee_manager, company_admin)

    if iotx_staking.hasRole(role_fee_manager, personal_admin):
        iotx_staking.renounceRole(role_fee_manager, personal_admin, {'from': personal_admin})
    assert not iotx_staking.hasRole(role_fee_manager, personal_admin)

    # Permission transferred:: ROLE_PAUSER (Personal Admin --> Company Admin)
    if not iotx_staking.hasRole(role_pauser, company_admin):
        iotx_staking.grantRole(role_pauser, company_admin, {'from': personal_admin})
    assert iotx_staking.hasRole(role_pauser, company_admin)

    if iotx_staking.hasRole(role_pauser, personal_admin):
        iotx_staking.renounceRole(role_pauser, personal_admin, {'from': personal_admin})
    assert not iotx_staking.hasRole(role_pauser, personal_admin)

    # Permission transferred:: DEFAULT_ADMIN_ROLE (Personal Admin --> Company Admin)
    if not iotx_staking.hasRole(role_default_admin, company_admin):
        iotx_staking.grantRole(role_default_admin, company_admin, {'from': personal_admin})
    assert iotx_staking.hasRole(role_default_admin, company_admin)

    if iotx_staking.hasRole(role_default_admin, personal_admin):
        iotx_staking.renounceRole(role_default_admin, personal_admin, {'from': personal_admin})
    assert not iotx_staking.hasRole(role_default_admin, personal_admin)
