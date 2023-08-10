# RockX Liquid Staking on [IoTeX](https://stake.IOTX.io/)
This project is currently under active development.

Current State: deployed at IoTeX testnet via [scripts/deploy/testnet](https://github.com/RockX-SG/uniiotx/blob/main/scripts/deploy/testnet.py).

Next State:  integration test on IoTeX testnet && unit test on local testnet.

### Network Configuration
To add IoTeX Testnet and Mainnet configurations, run the following commands separately:

- For IoTeX Testnet: `brownie networks add IoTeX iotex-testnet host=https://babel-api.testnet.iotex.io chainid=4690` 

- For IoTeX Mainnet: `brownie networks add IoTeX iotex-mainnet host=https://babel-api.mainnet.iotex.io chainid=4689`

If you successfully added the network, you'll see a success message along with the network details in the terminal.

To view the complete list of supported networks, you can run the following command: `brownie networks list`

![network_config](/docs/network_config.png) <br>

For further information, please visit the [Brownie Network Management](https://eth-brownie.readthedocs.io/en/stable/network-management.html#) 
and [IoTeX Ethereum API Compatibility](https://docs.iotex.io/reference/babel-web3-api) websites.

### Account Configuration
You can add a new account by running the following command: `brownie accounts new {INSERT-ACCOUNT-NAME}`

Make sure to replace {INSERT-ACCOUNT-NAME} with your name of choice. The following account names are used to
[deploy](https://github.com/RockX-SG/uniiotx/blob/main/scripts/deploy/testnet.py) 
and [upgrade](https://github.com/RockX-SG/uniiotx/tree/main/scripts/upgrade) contracts on IoTeX Testnet: 
- IoTeXAdmin 
- IoTeXDeployer 
- IoTeXOracle

You'll be prompted to enter in your private key and a password to encrypt the account with. If the account was configured successfully,
you'll see your account address printed to the terminal.

To view the complete list of accounts, you can run the following command: `brownie accounts list`

![accounts_list](/docs/accounts_list.png) <br>

For further information, please visit the [Brownie Working wih Accounts](https://eth-brownie.readthedocs.io/en/stable/core-accounts.html) websites.

### Unit Testing
Before running unit tests, programmers must have [Ganache](https://github.com/trufflesuite/ganache) installed.

All unit testing files can be found at [uniiotx/tests](https://github.com/RockX-SG/uniiotx/tree/main/tests).
Specifically:
- The file [uniiotx/tests/configs.py](https://github.com/RockX-SG/uniiotx/blob/main/tests/configs.py) contains all global variables, including various accounts.
- The file [uniiotx/tests/contracts.py](https://github.com/RockX-SG/uniiotx/blob/main/tests/contracts.py) includes a universal function for deploying all contracts.

Here are commonly used commands related to unit testing:

- To run the complete test suite: `brownie test`.

- To run a specific test: `brownie test tests/test_deposit.py`.

- To debug the project during test execution, add the `--interactive` flag: `brownie test --interactive`.

- To generate a gas profile report, add the `--gas` flag: `brownie test --gas`

- To evaluate unit test coverage, add the `--coverage` flag: `brownie test --coverage`

For further information, please visit the [Brownie Writing Unit Tests](https://eth-brownie.readthedocs.io/en/stable/tests-pytest-intro.html) websites.

**Todo:** Fix problems associated with executing the test suite. 

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

