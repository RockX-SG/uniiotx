# RockX Liquid Staking on [IoTeX](https://stake.IOTX.io/)
This project is currently under active development.

Current State: deployed at IoTeX testnet via [scripts/deploy/testnet](https://github.com/RockX-SG/uniiotx/blob/main/scripts/deploy/testnet.py).

Next State:  integration test on IoTeX testnet && unit test on local testnet.

### IoTeX Network Configurations
To add IoTeX Testnet and Mainnet configurations, execute the following commands separately:

For IoTeX Testnet: `brownie networks add IoTeX iotex-testnet host=https://babel-api.testnet.iotex.io chainid=4690` 

For IoTeX Mainnet: `brownie networks add IoTeX iotex-mainnet host=https://babel-api.mainnet.iotex.io chainid=4689`

If you successfully added the network, you'll see a success message along with the network details in the terminal.

To view the complete list of supported networks, you can run the following command: `brownie networks list`

![network_config](/docs/network_config.png) <br>

For further information, please visit the [Brownie Network Management](https://eth-brownie.readthedocs.io/en/stable/network-management.html#) 
and [IoTeX Ethereum API Compatibility](https://docs.iotex.io/reference/babel-web3-api) websites.


### Error Codes from Contracts
1. SYS001: INACTIVE_BUCKET_TYPE
1. SYS002: MANAGER_FEE_SHARES_OUT_OF_RANGE
1. SYS003: INVALID_DEBT_AMOUNT
1. SYS004: EMPTY_QUEUE
1. USR001: TRANSACTION_EXPIRED
1. USR002: INVALID_DEPOSIT_AMOUNT
1. USR003: INVALID_REDEEM_AMOUNT
1. USR004: INSUFFICIENT_ACCOUNTED_PRINCIPAL
1. USR005: INSUFFICIENT_ACCOUNTED_REWARD
1. USR006: INSUFFICIENT_ACCOUNTED_MANAGER_REWARD 
1. USR007: INVALID_TOKEN_AMOUNT_FOR_DEBT_PAYMENT
1. USR008: INVALID_TOTAL_PRINCIPAL_FOR_DEBT_PAYMENT

