from brownie import UniIOTX, IOTXStaking, accounts
from web3 import Web3
import threading

# The command to run this script: `brownie run scripts/test/stress_redeem.py  --network=iotex-mainnet`

# This script is utilized for the Stress_02 test suite, as outlined in the document
# "Scheme for Testing Bedrock Liquid Staking (IoTeX)". The document can be found at:
# https://github.com/RockX-SG/uniiotx/blob/main/docs/Scheme_for_Testing_Bedrock_Liquid_Staking%20(IoTeX).pdf

def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.mainnet.iotex.io'))

    # Contract
    uni_iotx = UniIOTX.at("0x236f8c0a61dA474dB21B693fB2ea7AAB0c803894")
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    # Account
    staker4 = accounts.load("Staker4")
    staker5 = accounts.load("Staker5")
    staker6 = accounts.load("Staker6")

    # Gas limit
    gas_limit = 6721975
    gas_price = 1000000000000

    # Parameters
    deadline = w3.eth.get_block('latest').timestamp + 3600
    amt_base = iotx_staking.getRedeemAmountBase()

    # Concurrent redeeming
    def redeem(address):
        while uni_iotx.balanceOf(address) * iotx_staking.exchangeRatio() / 1e18 >= amt_base:
            uni_iotx.approve(iotx_staking, amt_base, {'from': address, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
            tx = iotx_staking.redeem(amt_base, deadline, {'from': address, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
            assert "Redeemed" in tx.events

    thread1 = threading.Thread(target=redeem, args=([staker4]))
    thread2 = threading.Thread(target=redeem, args=([staker5]))
    thread3 = threading.Thread(target=redeem, args=([staker6]))

    thread1.start()
    thread2.start()
    thread3.start()

    thread1.join()
    thread2.join()
    thread3.join()

    print("Concurrent redeeming done")




