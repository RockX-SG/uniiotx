from brownie import IOTXStaking, accounts
from web3 import Web3
import time
import random
import threading

# The command to run this script: `brownie run scripts/test/stress_deposit.py  --network=iotex-mainnet`

# This script is utilized for the Stress_01 test suite, as outlined in the document
# "Scheme for Testing Bedrock Liquid Staking (IoTeX)". The document can be found at:
# https://github.com/RockX-SG/uniiotx/blob/main/docs/Scheme_for_Testing_Bedrock_Liquid_Staking%20(IoTeX).pdf

def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.mainnet.iotex.io'))

    # Contract
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    # Account
    staker4 = accounts.load("Staker4")
    staker5 = accounts.load("Staker5")
    staker6 = accounts.load("Staker6")

    # Gas limit
    gas_limit = 6721975
    gas_price = 1000000000000
    gas_fee_max_per_tx = gas_limit * gas_price

    # Parameters
    deadline = w3.eth.get_block('latest').timestamp + 3600
    stake_amt = 10000 * 1e18

    total_deposits = 100
    total_iotx = total_deposits * stake_amt

    # Check account balance
    total_gas_fee = gas_fee_max_per_tx * total_deposits
    min_balance_required = total_iotx + total_gas_fee
    for account in [staker4, staker5, staker6]:
        balance0 = account.balance()
        if balance0 <= min_balance_required:
            print("Account: ", account.address)
            print("Insufficient balance. Minium balance required: ", min_balance_required)
            print("You need to deposit extra IOTX into the account: ", min_balance_required - balance0)
            return

    print("Concurrent depositing begun...")

    # Concurrent depositing
    staked_tokens_2 = iotx_staking.getStakedTokenCount(2)

    def deposit(address):
        for i in range(total_deposits):
            print("Account: ", address)
            tx = iotx_staking.deposit(deadline, {'from': address, 'value': stake_amt, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
            assert "Minted" in tx.events
            assert "Staked" in tx.events
            sleep_time = random.uniform(1, 5)
            time.sleep(sleep_time)

    thread1 = threading.Thread(target=deposit, args=([staker4]))
    thread2 = threading.Thread(target=deposit, args=([staker5]))
    thread3 = threading.Thread(target=deposit, args=([staker6]))

    thread1.start()
    thread2.start()
    thread3.start()

    thread1.join()
    thread2.join()
    thread3.join()

    assert iotx_staking.getStakedTokenCount(2) == staked_tokens_2 + 3

    print("Concurrent depositing done")




