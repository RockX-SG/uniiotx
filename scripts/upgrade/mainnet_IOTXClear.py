from brownie import IOTXClear, accounts, project, config
from pathlib import Path

# Note: This script file is currently in draft form. Please do not execute it.
# Todo: Confirm the parameters again.

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/mainnet_IOTXClear.py  --network=iotex-mainnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'
    gas_price = '1000000000000'

    iotx_clear_proxy_addr = "0x7AD800771743F4e29f55235A55895273035FB546"
    iotx_clear_proxy = TransparentUpgradeableProxy.at(iotx_clear_proxy_addr)

    iotx_clear_upgraded = IOTXClear.deploy({'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})
    iotx_clear_proxy.upgradeTo(iotx_clear_upgraded, {'from': deployer, 'gas_limit': gas_limit, 'gas_price': gas_price})

    print("Upgraded IOTXClear address:", iotx_clear_upgraded)  # https://iotexscan.io/address/0x9cb04A403c6691AdEE14cE43D19aF1fE96C5ab91#transactions


