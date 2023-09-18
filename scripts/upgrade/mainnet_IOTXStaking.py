from brownie import IOTXStaking, accounts, project, config
from pathlib import Path

# Note: This script file is currently in draft form. Please do not execute it.
# Todo: Confirm the parameters again.

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/mainnet_IOTXStaking.py  --network=iotex-testnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'

    iotx_staking_proxy_addr = "0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f"
    iotx_staking_proxy = TransparentUpgradeableProxy.at(iotx_staking_proxy_addr)

    iotx_staking_upgraded = IOTXStaking.deploy({'from': deployer, 'gas_limit': gas_limit})
    iotx_staking_proxy.upgradeTo(iotx_staking_upgraded, {'from': deployer, 'gas_limit': gas_limit})

    print("Upgraded IOTXStaking address:", iotx_staking_upgraded)  # https://testnet.iotexscan.io/address/0x8783415300Fdaaf142db25A329FF21A294Bdd460#transactions


