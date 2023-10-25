from brownie import IOTXClear, accounts
from web3 import Web3
import threading

# The command to run this script: `brownie run scripts/test/stress_claim_mainnet_fork.py  --network=iotex-mainnet-fork`

# This script is utilized for the Stress_03 test suite, as outlined in the document
# "Scheme for Testing Bedrock Liquid Staking (IoTeX)". The document can be found at:
# https://github.com/RockX-SG/uniiotx/blob/main/docs/Scheme_for_Testing_Bedrock_Liquid_Staking%20(IoTeX).pdf

def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.mainnet.iotex.io'))

    # Contract
    iotx_clear = IOTXClear.at("0x7AD800771743F4e29f55235A55895273035FB546")

    # Account
    staker4 = accounts.at("0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09", True)
    staker5 = accounts.at("0x065e1164818487818E6BA714E8d80B91718ad758", True)
    staker6 = accounts.at("0xE88eBFccF58aaf553134AE5f87a77d0608B76d53", True)

    # Gas limit
    gas_limit = 6721975
    gas_price = 1000000000000

    # Concurrent claiming
    def claim(address):
        bal = iotx_clear.getBalance(address)
        if bal > 0:
            tx = iotx_clear.claim(bal, address, {'from': address, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
            assert "Claimed" in tx.events

    thread1 = threading.Thread(target=claim, args=([staker4]))
    thread2 = threading.Thread(target=claim, args=([staker5]))
    thread3 = threading.Thread(target=claim, args=([staker6]))

    thread1.start()
    thread2.start()
    thread3.start()

    thread1.join()
    thread2.join()
    thread3.join()

    print("Concurrent claiming done")




