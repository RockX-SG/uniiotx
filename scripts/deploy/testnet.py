from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts, Contract, project, config
from pathlib import Path
from web3 import Web3

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
    uni_iotx = UniIOTX.deploy({'from': deployer})  # https://testnet.iotexscan.io/address/0x472AC984B40D1e73Ac6ad3642bfbBE3be3D778E1#transactions
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer})  # https://testnet.iotexscan.io/address/0x73fA15Aa90F4F54F093bc7ECDF3906633D2eB194#transactions

    iotx_clear = IOTXClear.deploy({'from': deployer})  # https://testnet.iotexscan.io/address/0x54aE7559E11E4Ed483C9b5557E9aF349d1d29b1b#transactions
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer})  # https://testnet.iotexscan.io/address/0x703832231D3AbBDef4833cCC510ac7E06abDE5b7#transactions

    iotx_stake = IOTXStake.deploy({'from': deployer})  # https://testnet.iotexscan.io/address/0x6141572c66D5ff818678576a5790226047DF0bc0#transactions
    iotx_stake_proxy = TransparentUpgradeableProxy.deploy(iotx_stake, deployer, b'', {'from': deployer})  # https://testnet.iotexscan.io/address/0x22Ca5Bf95e7d8514FfE0605d3fFb1993660F3280#transactions

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_stake_transparent = Contract.from_abi("IOTXStake", iotx_stake_proxy.address, IOTXStake.abi)

    print("Deployed UniIOTX address:", uni_iotx_transparent)  # https://testnet.iotexscan.io/address/0x73fA15Aa90F4F54F093bc7ECDF3906633D2eB194#transactions
    print("Deployed IOTXClear address:", iotx_clear_transparent)  # https://testnet.iotexscan.io/address/0x703832231D3AbBDef4833cCC510ac7E06abDE5b7#transactions
    print("Deployed IOTXStake address:", iotx_stake_transparent)  # https://testnet.iotexscan.io/address/0x22Ca5Bf95e7d8514FfE0605d3fFb1993660F3280#transactions

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

    # Update ROLE_ORACLE provided by backend developer
    w3 = Web3(Web3.HTTPProvider('https://babel-api.testnet.iotex.io'))
    role_oracle = w3.keccak(text='ROLE_ORACLE')
    oracle_xie = "0x912AD2282799C5d62334017578418471f5aF5353"
    iotx_stake_transparent.grantRole(role_oracle, oracle_xie, {'from': admin})




