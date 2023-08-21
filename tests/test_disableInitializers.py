import brownie
from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStaking, accounts, Contract, project, config


def test_disableInitializers(proxy, fund, owner, deployer, admin, oracles, delegates, stake_duration, start_amount, common_ratio, sequence_length, stake_amounts, manager_fee_shares):
    # Deploy contracts
    system_staking = SystemStaking.deploy({'from': owner})

    uni_iotx = UniIOTX.deploy({'from': deployer})
    uni_iotx_proxy = proxy.deploy(uni_iotx, deployer, b'', {'from': deployer})

    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_clear_proxy = proxy.deploy(iotx_clear, deployer, b'', {'from': deployer})

    iotx_staking = IOTXStaking.deploy({'from': deployer})
    iotx_stake_proxy = proxy.deploy(iotx_staking, deployer, b'', {'from': deployer})

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_staking_transparent = Contract.from_abi("IOTXStaking", iotx_stake_proxy.address, IOTXStaking.abi)

    # The contracts are secured post-deployment to prevent their initialization or reinitialization to any version.
    for stake_amount in stake_amounts:
        system_staking.addBucketType(stake_amount, stake_duration, {'from': owner})

    with brownie .reverts("Initializable: contract is already initialized"):
        uni_iotx.initialize(iotx_staking_transparent, {'from': admin})
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_staking.initialize(
            system_staking,
            uni_iotx_transparent,
            iotx_clear_transparent,
            oracles[0],
            start_amount,
            common_ratio,
            sequence_length,
            stake_duration,
            {'from': admin}
        )
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_clear.initialize(system_staking, iotx_staking_transparent, oracles[0], {'from': admin})

    uni_iotx_transparent.initialize(iotx_staking_transparent, {'from': admin})
    iotx_staking_transparent.initialize(
        system_staking,
        uni_iotx_transparent,
        iotx_clear_transparent,
        oracles[0],
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin}
    )
    iotx_clear_transparent.initialize(system_staking, iotx_staking_transparent, oracles[0], {'from': admin})

    iotx_staking_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_staking_transparent.setGlobalDelegate(delegates[0], {'from': oracles[0]})

    with brownie .reverts("Initializable: contract is already initialized"):
        uni_iotx_transparent.initialize(iotx_staking, {'from': admin})
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_staking_transparent.initialize(
            system_staking,
            uni_iotx,
            iotx_clear,
            oracles[0],
            start_amount,
            common_ratio,
            sequence_length,
            stake_duration,
            {'from': admin}
        )
    with brownie .reverts("Initializable: contract is already initialized"):
        iotx_clear_transparent.initialize(system_staking, iotx_staking, oracles[0], {'from': admin})
