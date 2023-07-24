from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts, Contract, project, config
from pathlib import Path

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial
def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    # Bucket info
    start_amount = 100 * 1e18
    common_ratio = 10
    sequence_length = 3
    stake_duration = 34560

    # Manager fee shares
    manager_fee_shares = 100

    # SystemStaking contract address
    # https://testnet.iotexscan.io/address/0x52ab0fe2c3a94644de0888a3ba9ea1443672e61f#transactions
    system_staking = "0x52ab0fe2c3a94644de0888a3ba9ea1443672e61f"

    # Load accounts
    # Todo: Use a dedicated account, maybe consider ProxyAdmin contract
    deployer = accounts.load("IoTeXDeployer")
    admin = accounts.load("IoTeXAdmin")
    oracle = accounts.load("IoTeXOracle")

    # Fund accounts
    # Assume that the admin account holds sufficient assets
    fund_base = 1e18
    if deployer.balance() < fund_base:
        admin.transfer(deployer, fund_base*10)
    if oracle.balance() < fund_base:
        admin.transfer(oracle, fund_base*10)

    # Init delegate
    delegate = "0xac82586b613d10a33df00835aC6DAd8541952334"  # io14jp9s6mp85g2x00spq66cmdds4qe2ge5r0p72d

    # Deploy contracts
    uni_iotx = UniIOTX.deploy({'from': deployer})  # 0x2411AB8dE5abb1191F819e086728d37C35aA7EcB
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})  # 0x3423AC3e8E780C1081C5a1194E7f862fB1e09d5F

    iotx_clear = IOTXClear.deploy({'from': deployer})  # 0x7AD800771743F4e29f55235A55895273035FB546
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})  # 0x54B045860E49711eABDa160eBd5db8be1fC85A55

    iotx_stake = IOTXStake.deploy({'from': deployer})  # 0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f
    iotx_stake_proxy = TransparentUpgradeableProxy.deploy(iotx_stake, deployer, b'', {'from': deployer})  # 0x802d4900209b2292bF7f07ecAE187f836040A709

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_stake_transparent = Contract.from_abi("IOTXStake", iotx_stake_proxy.address, IOTXStake.abi)

    print("Deployed UniIOTX address:", uni_iotx_transparent)  # 0x3423AC3e8E780C1081C5a1194E7f862fB1e09d5F
    print("Deployed IOTXClear address:", iotx_clear_transparent)  # 0x54B045860E49711eABDa160eBd5db8be1fC85A55
    print("Deployed IOTXStake address:", iotx_stake_transparent)  # 0x802d4900209b2292bF7f07ecAE187f836040A709

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
    iotx_stake_transparent.setGlobalDelegate(delegate, {'from': oracle})





