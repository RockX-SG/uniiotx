from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStaking, accounts, Contract, project, config
from pathlib import Path
from web3 import Web3

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial

# The command to run this script: `brownie run scripts/deploy/testnet.py  --network=iotex-testnet`


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
    gas_limit = '6721975'
    uni_iotx = UniIOTX.deploy({'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0xa223567B74581F1B995b346e36e26C291d1cC1B4#transactions
    uni_iotx_proxy = TransparentUpgradeableProxy.deploy(uni_iotx, deployer, b'', {'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0x956a03ecEb344eA15A6CbE8949088992fAD88628#transactions

    iotx_clear = IOTXClear.deploy({'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0x7c193769f46D2B32819eb76E7c9Aeb580260A668#transactions
    iotx_clear_proxy = TransparentUpgradeableProxy.deploy(iotx_clear, deployer, b'', {'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0x4DC32Ad7BffAF50434b12195D3b59CD66601335D#transactions

    iotx_staking = IOTXStaking.deploy({'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0x171977bf57e93C9be280569F4853325719130C22#transactions
    iotx_staking_proxy = TransparentUpgradeableProxy.deploy(iotx_staking, deployer, b'', {'from': deployer, 'gas_limit': gas_limit})  # https://testnet.iotexscan.io/address/0xa479659F378d54168CD7859f5025133382EdB3C5#transactions

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_staking_transparent = Contract.from_abi("IOTXStaking", iotx_staking_proxy.address, IOTXStaking.abi)

    print("Deployed UniIOTX address:", uni_iotx_transparent)  # https://testnet.iotexscan.io/address/0x956a03ecEb344eA15A6CbE8949088992fAD88628#transactions
    print("Deployed IOTXClear address:", iotx_clear_transparent)  # https://testnet.iotexscan.io/address/0x4DC32Ad7BffAF50434b12195D3b59CD66601335D#transactions
    print("Deployed IOTXStaking address:", iotx_staking_transparent)  # https://testnet.iotexscan.io/address/0xa479659F378d54168CD7859f5025133382EdB3C5#transactions

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

    # Update ROLE_ORACLE provided by backend developer
    w3 = Web3(Web3.HTTPProvider('https://babel-api.testnet.iotex.io'))
    role_oracle = w3.keccak(text='ROLE_ORACLE')
    oracle_xie = "0x912AD2282799C5d62334017578418471f5aF5353"
    iotx_staking_transparent.grantRole(role_oracle, oracle_xie, {'from': admin})
    iotx_clear_transparent.grantRole((role_oracle, oracle_xie, {'from': admin}))




