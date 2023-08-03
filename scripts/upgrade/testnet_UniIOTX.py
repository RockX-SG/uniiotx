from brownie import UniIOTX, accounts, project, config
from pathlib import Path

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial

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

    print("Upgraded UniIOTX address:", uni_iotx_upgraded)  # https://testnet.iotexscan.io/address/0xEa9D88eBb1dBF8A0b41a531e95aa39779d82A9bf#transactions


