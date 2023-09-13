from brownie import UniIOTX, IOTXClear, IOTXStaking, accounts
from web3 import Web3

# The command to run this script: `brownie run scripts/authorize/admin_testnet_fork.py  --network=iotex-testnet-fork`

def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.testnet.iotex.io'))

    # Contract roles
    role_pauser = w3.keccak(text='ROLE_PAUSER')
    role_minter = w3.keccak(text='ROLE_MINTER')
    role_staker = w3.keccak(text='ROLE_STAKER')
    role_fee_manager = w3.keccak(text='ROLE_FEE_MANAGER')
    role_oracle = w3.keccak(text='ROLE_ORACLE')
    role_default_admin = w3.toBytes(hexstr="0x00")

    # Contracts, managed by the following personal Admin account
    contracts = [UniIOTX.at("0x956a03ecEb344eA15A6CbE8949088992fAD88628"),
                 IOTXClear.at("0x4DC32Ad7BffAF50434b12195D3b59CD66601335D"),
                 IOTXStaking.at("0xa479659F378d54168CD7859f5025133382EdB3C5")]

    # Personal Admin account
    admin_personal = accounts.at("0xbFdDf5e269C74157b157c7DaC5E416d44afB790d", True)

    # Fund account
    fund_base = 1e18
    accounts[0].transfer(admin_personal, fund_base*1000)

    # Fake company Admin account, which will take over the aforementioned contracts
    admin_company = "0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09"

    # Grant/Renounce DEFAULT_ADMIN_ROLE
    for contract in contracts:
        if contract.hasRole(role_default_admin, admin_personal):
            if not contract.hasRole(role_default_admin, admin_company):
                contract.grantRole(role_default_admin, admin_company, {'from': admin_personal})
            contract.renounceRole(role_default_admin, admin_personal, {'from': admin_personal})

    # Check Admin permission
    for contract in contracts:
        assert not contract.hasRole(role_default_admin, admin_personal)
        assert contract.hasRole(role_default_admin, admin_company)

    print("Company Admin granted: ", admin_company)
    print("Personal Admin renounced: ", admin_personal)
