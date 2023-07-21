from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts, Contract, project, config

# The configuration of the IoTeX testnet is necessary to run this script.
# Please refer to this link to obtain the chainID and endpoint information:
# https://docs.iotex.io/reference/babel-web3-api
# Todo: add a tutorial
def main():
    # Bucket info
    start_amount = 100 * 1e18
    common_ratio = 10
    sequence_length = 3
    stake_duration = 34560

    # Manager fee shares
    manager_fee_shares = 100

    # SystemStaking contract address
    # https://testnet.iotexscan.io/address/0x52ab0fe2c3a94644de0888a3ba9ea1443672e61f#transactions
    system_staking = "0x52ab0fe2c3a94644de0888a3ba9ea1443672e61f"

    # Oracle address
    oracle = "0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09"

    # Private accounts
    deployer = accounts.load("IoTeXDeployer")
    admin = accounts.load("IoTeXAdmin")

    # Init delegate
    delegate = "0xac82586b613d10a33df00835aC6DAd8541952334"  # io14jp9s6mp85g2x00spq66cmdds4qe2ge5r0p72d

    # Deploy contracts
    uni_iotx = UniIOTX.deploy({'from': deployer})
    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_stake = IOTXStake.deploy({'from': deployer})

    print("Deployed UniIOTX address:", uni_iotx)  # 0x85792f60633DBCF7c2414675bcC0a790B1b65CbB
    print("Deployed IOTXClear address:", iotx_clear)  # 0x06c186Ff3a0dA2ce668E5B703015f3134F4a88Ad
    print("Deployed IOTXStake address:", iotx_stake)  # 0x409dF0Ee22F057A88c1F98743c2e42D0c45DaDda

    uni_iotx.initialize(iotx_stake, {'from': admin})
    iotx_clear.initialize(system_staking, iotx_stake, oracle, {'from': admin})
    iotx_stake.initialize(
        system_staking,
        uni_iotx,
        iotx_clear,
        oracle,
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin}
    )
    iotx_stake.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_stake.setGlobalDelegate(delegate, {'from': oracle})





