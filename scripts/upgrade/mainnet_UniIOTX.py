from brownie import UniIOTX, accounts, project, config
from pathlib import Path

# Note: This script file is currently in draft form. Please do not execute it.
# Todo: Confirm the parameters again.

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/mainnet_UniIOTX.py  --network=iotex-mainnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'
    gas_price = '1000000000000'

    uni_iotx_proxy_addr = "0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894"
    uni_iotx_proxy = TransparentUpgradeableProxy.at(uni_iotx_proxy_addr)

    uni_iotx_upgraded = UniIOTX.deploy({'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})
    uni_iotx_proxy.upgradeTo(uni_iotx_upgraded, {'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})

    print("Upgraded UniIOTX address:", uni_iotx_upgraded)  # https://iotexscan.io/address/0x16221CaD160b441db008eF6DA2d3d89a32A05859#transactions


