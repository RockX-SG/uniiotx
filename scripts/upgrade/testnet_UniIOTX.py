from brownie import UniIOTX, accounts, project, config
from pathlib import Path

# Todo:
# Reminder: Please ensure to update the relevant addresses once the contracts have been successfully upgraded.

# The command to run this script: `brownie run scripts/upgrade/testnet_UniIOTX.py  --network=iotex-testnet`


def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy

    deployer = accounts.load("IoTeXDeployer")
    gas_limit = '6721975'

    uni_iotx_proxy_addr = "0x956a03ecEb344eA15A6CbE8949088992fAD88628"
    uni_iotx_proxy = TransparentUpgradeableProxy.at(uni_iotx_proxy_addr)

    uni_iotx_upgraded = UniIOTX.deploy({'from': deployer, 'gas_limit': gas_limit})
    uni_iotx_proxy.upgradeTo(uni_iotx_upgraded, {'from': deployer, 'gas_limit': gas_limit})

    print("Upgraded UniIOTX address:", uni_iotx_upgraded)  # https://testnet.iotexscan.io/address/0x209925387e9293cc2d901c684b569720D10c19eB#transactions


