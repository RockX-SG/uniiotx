from brownie import IOTXStaking, accounts, project, config
from pathlib import Path

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/testnet_IOTXStaking.py  --network=iotex-testnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'

    iotx_staking_proxy_addr = "0xa479659F378d54168CD7859f5025133382EdB3C5"
    iotx_staking_proxy = TransparentUpgradeableProxy.at(iotx_staking_proxy_addr)

    iotx_staking_upgraded = IOTXStaking.deploy({'from': deployer, 'gas_limit': gas_limit})
    iotx_staking_proxy.upgradeTo(iotx_staking_upgraded, {'from': deployer, 'gas_limit': gas_limit})

    print("Upgraded IOTXStaking address:", iotx_staking_upgraded)  # https://testnet.iotexscan.io/address/0x048a195c6bCDE90F086838a6f2C661FA53A534C6#transactions


