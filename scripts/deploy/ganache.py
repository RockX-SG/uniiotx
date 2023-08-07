from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStaking, accounts, Contract, project, config
from pathlib import Path
from brownie import *

# Note: The Ganache client must be installed locally prior to executing this script
# The link to the Ganache client: https://github.com/trufflesuite/ganache

# The command to run this script: `brownie run scripts/deploy/ganache.py`


def main():
    # Reset the local environment to the initial state
    chain.reset()

    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    # Bucket info
    start_amount = 100
    common_ratio = 10
    sequence_length = 3
    stake_duration = 1
    stake_amount0 = start_amount
    stake_amount1 = stake_amount0 * common_ratio
    stake_amount2 = stake_amount1 * common_ratio

    # Manager fee shares
    manager_fee_shares = 100

    # Prepare accounts
    # Todo: Use a dedicated account, maybe consider ProxyAdmin contract
    system_staking_owner = accounts[0]
    deployer = accounts[1]
    admin = accounts[2]
    oracle = accounts[3]
    delegate = accounts[4]
    user = accounts[5]


    # Deploy contracts
    system_staking = SystemStaking.deploy({'from': system_staking_owner})

    uni_iotx = UniIOTX.deploy({'from': deployer})
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})

    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})

    iotx_staking = IOTXStaking.deploy({'from': deployer})
    iotx_staking_proxy = TransparentUpgradeableProxy.deploy(iotx_staking, deployer, b'', {'from': deployer})

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_staking_transparent = Contract.from_abi("IOTXStaking", iotx_staking_proxy.address, IOTXStaking.abi)

    print("Deployed SystemStaking address:", system_staking)
    print("Deployed UniIOTX address:", uni_iotx_transparent)
    print("Deployed IOTXClear address:", iotx_clear_transparent)
    print("Deployed IOTXStaking address:", iotx_staking_transparent)

    # Configure contracts # Todo: Use loop for simplicity
    system_staking.addBucketType(stake_amount0, stake_duration, {'from': system_staking_owner})
    system_staking.addBucketType(stake_amount1, stake_duration, {'from': system_staking_owner})
    system_staking.addBucketType(stake_amount2, stake_duration, {'from': system_staking_owner})

    uni_iotx_transparent.initialize(iotx_staking_transparent, {'from': admin})
    iotx_staking_transparent.initialize(
        system_staking,
        uni_iotx_transparent,
        iotx_clear_transparent,
        oracle,
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin}
    )
    iotx_clear_transparent.initialize(system_staking, iotx_staking_transparent, oracle, {'from': admin})

    iotx_staking_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_staking_transparent.setGlobalDelegate(delegate, {'from': oracle})






