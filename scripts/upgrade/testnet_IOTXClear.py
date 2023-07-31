from brownie import IOTXClear, accounts, project, config
from pathlib import Path

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial

# The command to run this script: `brownie run scripts/upgrade/testnet_IOTXClear.py  --network=iotex-testnet`


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

    print("Upgraded IOTXClear address:", iotx_clear_upgraded)  # https://testnet.iotexscan.io/address/0x68Df31FB9b0229aBCbb3B9D30FB173Fdd51De995#transactions


