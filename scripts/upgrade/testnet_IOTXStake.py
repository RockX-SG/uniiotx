from brownie import IOTXStake, accounts, project, config
from pathlib import Path

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial

# The command to run this script: `brownie run scripts/upgrade/testnet_IOTXStake.py  --network=iotex-testnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'

    iotx_stake_proxy_addr = "0xa479659F378d54168CD7859f5025133382EdB3C5"
    iotx_stake_proxy = TransparentUpgradeableProxy.at(iotx_stake_proxy_addr)

    iotx_stake_upgraded = IOTXStake.deploy({'from': deployer, 'gas_limit': gas_limit})
    iotx_stake_proxy.upgradeTo(iotx_stake_upgraded, {'from': deployer, 'gas_limit': gas_limit})

    print("Upgraded IOTXStake address:", iotx_stake_upgraded)  # https://testnet.iotexscan.io/address/0x6378a9C643Cd85C0745860592335bab6740Ca517#transactions


