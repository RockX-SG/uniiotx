from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStaking, accounts, Contract, project, config
from pathlib import Path

# The command to run this script: `brownie run scripts/deploy/mainnet_fork.py  --network=iotex-mainnet-fork`

def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    # Bucket info
    start_amount = 10000 * 1e18
    common_ratio = 10
    sequence_length = 3
    stake_duration = 1572480  # (91 * 24 * 60 * 60) / 5

    # Manager fee shares
    manager_fee_shares = 50

    # SystemStaking contract address
    # https://iotexscan.io/address/0x68db92a6a78a39dcaff1745da9e89e230ef49d3d#code
    system_staking = "0x68db92a6a78a39dcaff1745da9e89e230ef49d3d"

    # Load accounts
    # Todo: Use a dedicated account, maybe consider ProxyAdmin contract
    deployer = accounts.at("0x3af43AfEd31C00311381A8DF26cc58C9bD2b5375", True)
    admin = accounts.at("0xbFdDf5e269C74157b157c7DaC5E416d44afB790d", True)
    oracle = accounts.at("0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09", True)

    # Fund accounts
    fund_base = 1e18
    accounts[0].transfer(deployer, fund_base*1000)
    accounts[0].transfer(admin, fund_base*1000)
    accounts[0].transfer(oracle, fund_base*1000)

    # Record balances before deployment
    deployer_bal_0 = deployer.balance()
    admin_bal_0 = admin.balance()
    oracle_bal_0 = oracle.balance()

    # Init delegate
    delegate = "0xFE82234dE6b7F4DAE188552dfAa64C1dF5674caA"

    # Deploy contracts
    gas_limit = '6721975'
    gas_price = 'auto'
    uni_iotx = UniIOTX.deploy({'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})

    iotx_clear = IOTXClear.deploy({'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})

    iotx_staking = IOTXStaking.deploy({'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})
    iotx_staking_proxy = TransparentUpgradeableProxy.deploy(iotx_staking, deployer, b'', {'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_staking_transparent = Contract.from_abi("IOTXStaking", iotx_staking_proxy.address, IOTXStaking.abi)

    print("Deployed UniIOTX address:", uni_iotx_transparent)
    print("Deployed IOTXClear address:", iotx_clear_transparent)
    print("Deployed IOTXStaking address:", iotx_staking_transparent)

    uni_iotx_transparent.initialize(iotx_staking_transparent, {'from': admin, 'gas_limit': gas_limit, 'gas_price': gas_price})
    iotx_staking_transparent.initialize(
        system_staking,
        uni_iotx_transparent,
        iotx_clear_transparent,
        oracle,
        admin,
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin, 'gas_limit': gas_limit, 'gas_price': gas_price}
    )
    iotx_clear_transparent.initialize(system_staking, iotx_staking_transparent, oracle, admin, {'from': admin, 'gas_limit': gas_limit, 'gas_price': gas_price})

    iotx_staking_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin, 'gas_limit': gas_limit, 'gas_price': gas_price})
    iotx_staking_transparent.setGlobalDelegate(delegate, {'from': oracle, 'gas_limit': gas_limit, 'gas_price': gas_price})

    # Deployment gas fee estimation
    deployer_bal_1 = deployer.balance()
    admin_bal_1 = admin.balance()
    oracle_bal_1 = oracle.balance()

    gas_fee_deployer = deployer_bal_0 - deployer_bal_1
    gas_fee_admin = admin_bal_0 - admin_bal_1
    gas_fee_oracle = oracle_bal_0 - oracle_bal_1

    print("Gas fee for Deployer: ", gas_fee_deployer)
    print("Gas fee for Admin: ", gas_fee_admin)
    print("Gas fee for Oracle: ", gas_fee_oracle)

    print("Total gas fee: ", gas_fee_deployer + gas_fee_admin + gas_fee_oracle)



