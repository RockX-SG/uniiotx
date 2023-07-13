from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts, Contract, project, config
from pathlib import Path


# DefaultAdmin address: "0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC"
# DefaultAdmin private key: "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

def main():
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    # Bucket info
    start_amount = 10000
    common_ratio = 10
    sequence_length = 3

    # Prepare accounts
    # Todo: Tune it properly
    deployer = accounts.load('DefaultAdmin')
    owner = accounts.load('Alice')
    orale = deployer

    # Deploy contracts
    system_staking = SystemStaking.deploy({'from': deployer})
    system_staking_proxy = TransparentUpgradeableProxy.deploy(system_staking, deployer, b'', {'from': deployer})
    system_staking_transparent = Contract.from_abi("SystemStaking", system_staking_proxy.address, SystemStaking.abi)

    uni_iotx = UniIOTX.deploy({'from': deployer})
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})
    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)

    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)

    iotx_stake = IOTXStake.deploy({'from': deployer})
    iotx_stake_proxy = TransparentUpgradeableProxy.deploy(iotx_stake, deployer, b'', {'from': deployer})
    iotx_stake_transparent = Contract.from_abi("IOTXStake", iotx_stake_proxy.address, IOTXStake.abi)

    print("SystemStaking address:", system_staking_transparent)
    print("UniIOTX address:", uni_iotx_transparent)
    print("IOTXClear address:", iotx_clear_transparent)
    print("IOTXStake address:", iotx_stake_transparent)

    # Configure contracts
    uni_iotx_transparent.initialize(iotx_stake_transparent, {'from': owner})
    iotx_clear_transparent.initialize(system_staking_transparent, iotx_stake_transparent, orale, {'from': owner})
    iotx_stake_transparent.initialize(
        system_staking_transparent,
        iotx_clear_transparent,
        orale,
        start_amount,
        common_ratio,
        sequence_length,
        {'from': owner}
    )






