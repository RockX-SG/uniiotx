from brownie import accounts, project, config
from pathlib import Path

# The command to run this script: `brownie run scripts/authorize/mainnet_proxy_admin.py  --network=iotex-mainnet`

# Executing the following scripts will:
# 1. Change the admins of UniIOTX, IOTXClear, IOTXStaking transparent proxies to the given ProxyAdmin contract.
# 2. Change the owner of given ProxyAdmin contract to company owner.

def main():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    # Reference: https://docs.openzeppelin.com/contracts/3.x/api/proxy#ProxyAdmin
    # Reference: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/transparent/ProxyAdmin.sol
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    TransparentUpgradeableProxy = deps.TransparentUpgradeableProxy
    ProxyAdmin = deps.ProxyAdmin

    # Contracts
    proxy_admin_contract = ProxyAdmin.at("0xb1e2a647A668F349b2D204E6058F41cbD8c6B9B6")

    uni_iotx_transparent_proxy_contract = TransparentUpgradeableProxy.at("0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894")
    iotx_clear_transparent_proxy_contract = TransparentUpgradeableProxy.at("0x7AD800771743F4e29f55235A55895273035FB546")
    iotx_staking_transparent_proxy_contract = TransparentUpgradeableProxy.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    transparent_proxies = [uni_iotx_transparent_proxy_contract, iotx_clear_transparent_proxy_contract, iotx_staking_transparent_proxy_contract]

    # Personal accounts
    personal_owner = accounts.load("IoTeXAdmin")
    personal_deployer = accounts.load("IoTeXDeployer")

    # Company account
    company_owner = "0xE20aC9B2889EF2e4d9B6A4678e9be8d6048861B8"

    # Transfer authority
    gas_limit = '6721975'
    gas_price = '1000000000000'

    # Permission to be transferred: Transparent Proxy Contract Admin (Personal Deployer --> ProxyAdmin Contract)
    for transparent_proxy in transparent_proxies:
        transparent_proxy.changeAdmin(proxy_admin_contract, {'from': personal_deployer, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
        assert proxy_admin_contract.getProxyAdmin(transparent_proxy) == proxy_admin_contract.address

    # Permission transferred:: ProxyAdmin Contract Owner (Personal owner --> Company owner)
    if proxy_admin_contract.owner() == personal_owner:
        proxy_admin_contract.transferOwnership(company_owner, {'from': personal_owner, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
    assert proxy_admin_contract.owner() == company_owner