import pytest
from web3 import Web3
from pathlib import Path
from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStaking, accounts, Contract, project, config

# Web3 client
@pytest.fixture(scope="session", autouse=True)
def w3():
    return Web3(Web3.HTTPProvider('http://localhost:8545'))

# Tx deadline
@pytest.fixture(scope="session", autouse=True)
def deadline(w3):
    return w3.eth.get_block('latest').timestamp + 3600

@pytest.fixture(scope="session", autouse=True)
def zero_address():
    return "0x0000000000000000000000000000000000000000"

# Roles
@pytest.fixture(scope="session", autouse=True)
def roles(w3):
    role_pause = w3.keccak(text='ROLE_PAUSER')  # index = 0
    role_mint = w3.keccak(text='ROLE_MINTER')  # index = 1
    role_stake = w3.keccak(text='ROLE_STAKER')  # index = 2
    role_fee_manager = w3.keccak(text='ROLE_FEE_MANAGER')  # index = 3
    role_oracle = w3.keccak(text='ROLE_ORACLE')  # index = 4
    role_default_admin = w3.toBytes(hexstr="0x00")  # index = 5
    return [role_pause, role_mint, role_stake, role_fee_manager, role_oracle, role_default_admin]

# Predefined Accounts
@pytest.fixture(scope="session", autouse=True)
def owner():
    return accounts[0]

@pytest.fixture(scope="session", autouse=True)
def deployer():
    return accounts[1]

@pytest.fixture(scope="session", autouse=True)
def admin():
    return accounts[2]

@pytest.fixture(scope="session", autouse=True)
def oracle():
    return accounts[3]

@pytest.fixture(scope="session", autouse=True)
def delegates():
    return [accounts[4], accounts[5]]

@pytest.fixture(scope="session", autouse=True)
def users():
    return [accounts[6], accounts[7]]

# Protocol params
@pytest.fixture(scope="session", autouse=True)
def start_amount():
    return 100

@pytest.fixture(scope="session", autouse=True)
def common_ratio():
    return 10

@pytest.fixture(scope="session", autouse=True)
def sequence_length():
    return 3

@pytest.fixture(scope="session", autouse=True)
def stake_duration():
    return 1

@pytest.fixture(scope="session", autouse=True)
def stake_amounts(start_amount, common_ratio):
    return [start_amount, start_amount * common_ratio, start_amount * common_ratio * common_ratio]

@pytest.fixture(scope="session", autouse=True)
def manager_fee_shares():
    return 100

@pytest.fixture(scope="session", autouse=True)
def proxy():
    # Reference: https://docs.openzeppelin.com/contracts/4.x/api/proxy#TransparentUpgradeableProxy
    deps = project.load(  Path.home() / ".brownie" / "packages" / config["dependencies"][0])
    return deps.TransparentUpgradeableProxy

# Contracts
# Note: The Ganache client must be installed locally prior to executing this script
# The link to the Ganache client: https://github.com/trufflesuite/ganache
@pytest.fixture()
def contracts(proxy, owner, deployer, admin, oracle, delegates, stake_duration, start_amount, common_ratio, sequence_length, stake_amounts, manager_fee_shares):
    # Deploy contracts
    system_staking = SystemStaking.deploy({'from': owner})

    uni_iotx = UniIOTX.deploy({'from': deployer})
    uni_iotx_proxy = proxy.deploy(uni_iotx, deployer, b'', {'from': deployer})

    iotx_clear = IOTXClear.deploy({'from': deployer})
    iotx_clear_proxy = proxy.deploy(iotx_clear, deployer, b'', {'from': deployer})

    iotx_staking = IOTXStaking.deploy({'from': deployer})
    iotx_stake_proxy = proxy.deploy(iotx_staking, deployer, b'', {'from': deployer})

    uni_iotx_transparent = Contract.from_abi("UniIOTX", uni_iotx_proxy.address, UniIOTX.abi)
    iotx_clear_transparent = Contract.from_abi("IOTXClear", iotx_clear_proxy.address, IOTXClear.abi)
    iotx_staking_transparent = Contract.from_abi("IOTXStaking", iotx_stake_proxy.address, IOTXStaking.abi)

    # Configure contracts
    for stake_amount in stake_amounts:
        system_staking.addBucketType(stake_amount, stake_duration, {'from': owner})

    uni_iotx_transparent.initialize(iotx_staking_transparent, {'from': admin})
    iotx_staking_transparent.initialize(
        system_staking,
        uni_iotx_transparent,
        iotx_clear_transparent,
        oracle,
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin}
    )
    iotx_clear_transparent.initialize(system_staking, iotx_staking_transparent, oracle, {'from': admin})

    iotx_staking_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_staking_transparent.setGlobalDelegate(delegates[0], {'from': oracle})

    return [system_staking, uni_iotx_transparent, iotx_clear_transparent, iotx_staking_transparent]
