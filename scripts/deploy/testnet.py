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

    # Oracle address
    oracle = "0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09"

    # Private accounts
    # Todo: Use a dedicated account, maybe consider ProxyAdmin contract
    deployer = accounts.load("IoTeXDeployer")
    admin = accounts.load("IoTeXAdmin")

    # Init delegate
    init_delegate = "0xac82586b613d10a33df00835aC6DAd8541952334"  # io14jp9s6mp85g2x00spq66cmdds4qe2ge5r0p72d

    # Deploy contracts
    uni_iotx = UniIOTX.deploy({'from': deployer})  # 0x57518941854F879f80fA1ABF2366f54339DE8436
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})  # 0x16221CaD160b441db008eF6DA2d3d89a32A05859
    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)

    iotx_clear = IOTXClear.deploy({'from': deployer})  # 0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})  # 0x97e16DB82E089D0C9c37bc07F23FcE98cfF04823
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)

    iotx_stake = IOTXStake.deploy({'from': deployer})  # 0x2ac98DB41Cbd3172CB7B8FD8A8Ab3b91cFe45dCf
    iotx_stake_proxy = TransparentUpgradeableProxy.deploy(iotx_stake, deployer, b'', {'from': deployer})  # 0x22B827e459654abd38ce48bF72d61c7372CD17c9
    iotx_stake_transparent = Contract.from_abi("IOTXStake", iotx_stake_proxy.address, IOTXStake.abi)

    print("Deployed UniIOTX address:", uni_iotx_transparent)  # 0x16221CaD160b441db008eF6DA2d3d89a32A05859
    print("Deployed IOTXClear address:", iotx_clear_transparent)  # 0x97e16DB82E089D0C9c37bc07F23FcE98cfF04823
    print("Deployed IOTXStake address:", iotx_stake_transparent)  # 0x22B827e459654abd38ce48bF72d61c7372CD17c9

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
    iotx_stake_transparent.setGlobalDelegate(init_delegate, {'from': oracle})
    iotx_stake_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})





