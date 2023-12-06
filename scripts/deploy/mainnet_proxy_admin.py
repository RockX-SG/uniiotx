from brownie import accounts, project, config
from pathlib import Path

# The command to run this script: `brownie run scripts/deploy/mainnet_proxy_admin.py  --network=iotex-mainnet`

def main():
    #  Reference: https://docs.openzeppelin.com/contracts/3.x/api/proxy#ProxyAdmin
    # Reference: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/proxy/transparent/ProxyAdmin.sol
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    ProxyAdmin = deps.ProxyAdmin

    # Load accounts
    owner = accounts.load("IoTeXAdmin")

    # Deploy contracts
    gas_limit = '6721975'
    gas_price = '1000000000000'

    proxyAdmin = ProxyAdmin.deploy({'from': owner, 'gas_limit': gas_limit, 'gas_price': gas_price})

    print("Deployed proxyAdmin address:", proxyAdmin)  # https://iotexscan.io/address/0xb1e2a647A668F349b2D204E6058F41cbD8c6B9B6#transactions



