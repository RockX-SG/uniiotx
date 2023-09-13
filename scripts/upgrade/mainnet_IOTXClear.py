from brownie import IOTXClear, accounts, project, config
from pathlib import Path

# Note: This script file is currently in draft form. Please do not execute it.
# Todo: Confirm the parameters again.

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/mainnet_IOTXClear.py  --network=iotex-testnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'

    iotx_clear_proxy_addr = "0x4DC32Ad7BffAF50434b12195D3b59CD66601335D"
    iotx_clear_proxy = TransparentUpgradeableProxy.at(iotx_clear_proxy_addr)

    iotx_clear_upgraded = IOTXClear.deploy({'from': deployer, 'gas_limit': gas_limit})
    iotx_clear_proxy.upgradeTo(iotx_clear_upgraded, {'from': deployer, 'gas_limit': gas_limit})

    print("Upgraded IOTXClear address:", iotx_clear_upgraded)  # https://testnet.iotexscan.io/address/0x18b05a0f14716DD7A8462EcdF9f0F009390BFf97#transactions

