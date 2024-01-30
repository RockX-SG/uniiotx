from brownie import IOTXStaking

# The command to run this script: `brownie run scripts/query/manager_reward_uniiotx.py  --network=iotex-mainnet`

def main():
    iotx_staking = IOTXStaking.at("0x2c914Ba874D94090Ba0E6F56790bb8Eb6D4C7e5f")
    reward = iotx_staking.getManagerReward()
    exchangeRatio = iotx_staking.exchangeRatio()
    print("Manager rewards (uniIOTX):", reward * 1e18 / exchangeRatio)