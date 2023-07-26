import pytest
from pathlib import Path
from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts, Contract, project, config
from configs import *
from brownie import *

# Contracts
# Note: The Ganache client must be installed locally prior to executing this script
# The link to the Ganache client: https://github.com/trufflesuite/ganache
@pytest.fixture
def contracts(owner, deployer, admin, oracle, delegates, stake_duration, start_amount, common_ratio, sequence_length, stake_amounts, manager_fee_shares):
    # Reset the local environment to the initial state
    chain.reset()

    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    # Deploy contracts
    system_staking = SystemStaking.deploy({'from': owner})

    uni_iotx = UniIOTX.deploy({'from': deployer})
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})

    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})

    iotx_stake = IOTXStake.deploy({'from': deployer})
    iotx_stake_proxy = TransparentUpgradeableProxy.deploy(iotx_stake, deployer, b'', {'from': deployer})

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_stake_transparent = Contract.from_abi("IOTXStake", iotx_stake_proxy.address, IOTXStake.abi)

    # Configure contracts
    for stake_amount in stake_amounts:
        system_staking.addBucketType(stake_amount, stake_duration, {'from': owner})

    uni_iotx_transparent.initialize(iotx_stake_transparent, {'from': admin})
    iotx_clear_transparent.initialize(system_staking, iotx_stake_transparent, oracle, {'from': admin})
    iotx_stake_transparent.initialize(
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
    iotx_stake_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_stake_transparent.setGlobalDelegate(delegates[0], {'from': oracle})

    return [system_staking, uni_iotx_transparent, iotx_clear_transparent, iotx_stake_transparent]