from brownie import IOTXStaking, accounts
from web3 import Web3

# The command to run this script: `brownie run scripts/test/merge.py  --network=iotex-mainnet`

# This script is utilized for the S1_T02 test suite, as outlined in the document
# "Scheme for Testing Bedrock Liquid Staking (IoTeX)". The document can be found at:
# https://github.com/RockX-SG/uniiotx/blob/main/docs/Scheme_for_Testing_Bedrock_Liquid_Staking%20(IoTeX).pdf

def main():
    # Web3 client
    w3 = Web3(Web3.HTTPProvider('https://babel-api.mainnet.iotex.io'))

    # Contract
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")

    # Account
    staker2 = accounts.load("Staker2")

    # Gas limit
    gas_limit = 6721975
    gas_price = 1000000000000
    gas_fee_max_per_tx = gas_limit * gas_price

    # Parameters
    deadline = w3.eth.get_block('latest').timestamp + 3600
    common_ratio = 10
    stake_amt_0 = 10000 * 1e18
    stake_amt_1 = stake_amt_0 * common_ratio

    # Total deposits at level 1
    staked_tokens_0 = iotx_staking.getStakedTokenCount(0)
    total_deposits_0 = common_ratio - staked_tokens_0
    total_iotx_0 = total_deposits_0 * stake_amt_0

    # Total deposits at level 2
    staked_tokens_1 = iotx_staking.getStakedTokenCount(1)
    total_deposits_1 = common_ratio - staked_tokens_1 - 1
    total_iotx_1 = total_deposits_1 * stake_amt_1

    # Check account balance
    total_gas_fee = gas_fee_max_per_tx * (total_deposits_0 + total_deposits_1)
    total_iotx = total_iotx_0 + total_iotx_1
    min_balance_required = total_iotx + total_gas_fee
    balance0 = staker2.balance()
    if balance0 <= min_balance_required:
        print("Insufficient balance. Minium balance required: ", min_balance_required)
        print("You need to deposit extra IOTX into the account: ", min_balance_required - balance0)
        return

    # Deposit the minimum amount in the bucket, resulting in a single merge operation at level 1.
    print("Begin depositing and merging at level 1...")
    for i in range(total_deposits_0 - 1):
        tx = iotx_staking.deposit(deadline, {'from': staker2, 'value': stake_amt_0, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
        assert "Minted" in tx.events
        assert "Staked" in tx.events
        assert "Merged" not in tx.events
    tx = iotx_staking.deposit(deadline, {'from': staker2, 'value': stake_amt_0, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert "Merged" in tx.events
    assert iotx_staking.getStakedTokenCount(0) == 0

    # Deposit the mid-sized amount in the bucket, resulting in a single merge operation at level 2.
    print("Begin depositing and merging at level 2...")
    staked_tokens_2 = iotx_staking.getStakedTokenCount(2)
    for i in range(total_deposits_1 - 1):
        tx = iotx_staking.deposit(deadline, {'from': staker2, 'value': stake_amt_1, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
        assert "Minted" in tx.events
        assert "Staked" in tx.events
        assert "Merged" not in tx.events
    tx = iotx_staking.deposit(deadline, {'from': staker2, 'value': stake_amt_1, 'gas_limit': gas_limit, 'gas_price': gas_price, 'allow_revert': True})
    assert "Minted" in tx.events
    assert "Staked" in tx.events
    assert "Merged" in tx.events
    assert iotx_staking.getStakedTokenCount(1) == 0
    assert iotx_staking.getStakedTokenCount(2) == staked_tokens_2 + 1

    print("Depositing and merging done. Total deposited IOTX: ", total_iotx)




