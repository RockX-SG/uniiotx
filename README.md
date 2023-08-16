# RockX Liquid Staking on [IoTeX](https://stake.IOTX.io/)
This project is currently under active development.

Current State: deployed at IoTeX testnet via [scripts/deploy/testnet](https://github.com/RockX-SG/uniiotx/blob/main/scripts/deploy/testnet.py).

Next State:  integration test on IoTeX testnet && unit test on local testnet.

### Prerequisites
This project, which includes [Solidity](https://soliditylang.org/) contracts, was developed using [Brownie](https://github.com/eth-brownie/brownie) on the [Pop!_OS](https://pop.system76.com/).

It also relies on the following projects:
[Python](https://www.python.org/), [Golang](https://go.dev/), [Ganache](https://github.com/trufflesuite/ganache), [Node.js](https://nodejs.org/en), [NPM](https://www.npmjs.com/), [NVM](https://github.com/nvm-sh/nvm), 
[Abigen](https://github.com/ethereum/go-ethereum/tree/master/cmd/abigen)

These are the currently adopted versions:
- Pop!_OS: v22.04 LTS
- Solidity: v0.8.9
- Python: v3.10.6
- Golang: v1.20.4
- Brownie: v1.19.3
- Ganache: v7.9.0
- Node.js: v20.5.0
- NPM: v9.8.1
- NVM: v0.39.0
- Abigen: v1.10.17

We assume that you have installed the necessary software on your development operating system other than 
Ganache, Node.js, NPM, NVM and Abigen.

#### Ganache Installation
Firstly, you need to install [Node.js](https://nodejs.org/en) and [NPM](https://www.npmjs.com/). We highly recommend using
[NVM](https://github.com/nvm-sh/nvm) for the installation of these two dependencies. 

You can run the following commands to install NVM, Node.js and NPM step by step:
- To install the latest version of NVM: `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash`<br>
After the installation, please restart your terminal or run the command: `source ~/.bashrc` for the changes to take effect.
- To install the latest version of Node.js: `nvm install node`
- To upgrade NPM to the latest version: `npm install -g npm@latest`

You can run the subsequent command to verify the version you've recently installed:
  - `nvm --version`
  - `node --version`
  - `npm --version`

Secondly, you can run the following command to install Ganache globally: `npm install ganache --global`.

Once installed globally, you can start ganache right from your command line: `ganache`

For further information, please visit the [Installing and Updating NVM](https://github.com/nvm-sh/nvm#installing-and-updating), 
[Installing Node.js and NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) and 
[Ganache](https://github.com/trufflesuite/ganache) websites 

#### Abigen Installation
We use Abigen to creates ABIs for Go native applications. You can run the following commands for installation:
- To download the source package: `wget -O $HOME/go-ethereum-1.10.17.tar.gz https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.10.17.tar.gz`
- To extract the archive: `cd $HOME && tar -C $HOME -xzf go-ethereum-1.10.17.tar.gz`
- To build the Abigen program: `cd $HOME/go-ethereum-1.10.17/cmd/abigen && go build .`
- To add the executable binary to the PATH environment variable: `export PATH=$PATH:$HOME/go-ethereum-1.10.17/cmd/abigen`

After the installation, please restart your terminal or run the command: `source ~/.bashrc` for the changes to take effect.

### Network Configuration
To add IoTeX Testnet and Mainnet configurations, run the following commands separately:

- For the live IoTeX Testnet: `brownie networks add IoTeX iotex-testnet host=https://babel-api.testnet.iotex.io chainid=4690`
- For the live IoTeX Mainnet: `brownie networks add IoTeX iotex-mainnet host=https://babel-api.mainnet.iotex.io chainid=4689`
- For the forked development IoTeX Testnet: `brownie networks add Development iotex-testnet-fork name="Ganache-CLI (IoTeX Testnet Fork)" host=http://127.0.0.1:8545 cmd=ganache fork=iotex-testnet `
- For the forked development IoTeX Mainnet: `brownie networks add Development iotex-mainnet-fork name="Ganache-CLI (IoTeX Mainnet Fork)" host=http://127.0.0.1:8545 cmd=ganache fork=iotex-mainnet`

Please be aware that forked development networks rely on the configurations of Ganache and live networks.

If you successfully added the network, you'll see a success message along with the network details in the terminal.

To view the complete list of supported networks, you can run the following command: `brownie networks list`

![network_config](/docs/network_config.png) <br>

For further information, please visit the [Brownie Network Management](https://eth-brownie.readthedocs.io/en/stable/network-management.html#) 
and [IoTeX Ethereum API Compatibility](https://docs.iotex.io/reference/babel-web3-api) websites.

### Account Configuration

You can add a new account by running the following command: `brownie accounts new {INSERT-ACCOUNT-NAME}`

Make sure to replace {INSERT-ACCOUNT-NAME} with your name of choice. The following account are currently used to
[deploy](https://github.com/RockX-SG/uniiotx/blob/main/scripts/deploy/testnet.py) 
and [upgrade](https://github.com/RockX-SG/uniiotx/tree/main/scripts/upgrade) contracts on IoTeX Testnet: 
- IoTeXAdmin: 0xbFdDf5e269C74157b157c7DaC5E416d44afB790d
- IoTeXDeployer: 0x3af43AfEd31C00311381A8DF26cc58C9bD2b5375
- IoTeXOracle: 0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09

You'll be prompted to enter in your private key and a password to encrypt the account with. If the account was configured successfully,
you'll see your account address printed to the terminal.

To view the complete list of accounts, you can run the following command: `brownie accounts list`

![accounts_list](/docs/accounts_list.png) <br>

Additionally, the accounts listed in the [configuration file](https://github.com/RockX-SG/uniiotx/blob/main/brownie-config.yaml)
are unlocked on the local Ganache and forked IoTeX Testnet for testing purposes:
- IoTeXSystemStakingOwner: 0x065e1164818487818E6BA714E8d80B91718ad758
- IoTeXAdmin: 0xbFdDf5e269C74157b157c7DaC5E416d44afB790d
- IoTeXDeployer: 0x3af43AfEd31C00311381A8DF26cc58C9bD2b5375
- IoTeXOracle1: 0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09
- IoTeXOracle2: 0x912AD2282799C5d62334017578418471f5aF5353
- IoTeXDelegate1: 0xac82586b613d10a33df00835aC6DAd8541952334  
- IoTexDelegate2: 0xE88eBFccF58aaf553134AE5f87a77d0608B76d53
- IoTeXUser1: 0x9ACE9968545089893208f35A81569Fa81cd24F7c
- IoTeXUser2: 0x912AD2282799C5d62334017578418471f5aF5353

For further information, please visit the [Brownie Working wih Accounts](https://eth-brownie.readthedocs.io/en/stable/core-accounts.html)
and [Brownie Unlocking Accounts on Development Networks](https://eth-brownie.readthedocs.io/en/stable/account-management.html#unlocking-accounts-on-development-networks) websites.

### Unit Testing
All the files for unit testing can be found in the [tests](https://github.com/RockX-SG/uniiotx/tree/main/tests) directory.

Specifically, the [tests/conftest.py](https://github.com/RockX-SG/uniiotx/blob/main/tests/conftest.py) file 
contains all global [fixtures](https://eth-brownie.readthedocs.io/en/stable/tests-pytest-intro.html#fixtures),
including various accounts. These fixtures are applied across multiple modules.

Here are commonly used commands related to unit testing:

- To run the complete test suite: `brownie test`.
- To run a specific test: `brownie test tests/test_deposit.py`.
- To only run updated tests, add the `--update` flag: `brownie test --update`.
- To debug the project during test execution, add the `--interactive` flag: `brownie test --interactive`.
- To generate a gas profile report, add the `--gas` flag: `brownie test --gas`
- To evaluate unit test coverage, add the `--coverage` flag: `brownie test --coverage`

The commands above operate on the local Ganache network. To conduct tests on the forked IoTeX Testnet, 
add the `--network iotex-testnet-fork` flag. For example: `brownie test --network iotex-testnet-fork`.

For further information, please visit the [Brownie Writing Unit Tests](https://eth-brownie.readthedocs.io/en/stable/tests-pytest-intro.html) websites.

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

