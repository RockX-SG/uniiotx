import pytest
from brownie import accounts
from web3 import Web3

# Roles
@pytest.fixture
def roles():
    # Update ROLE_ORACLE provided by backend developer
    w3 = Web3(Web3.HTTPProvider('http://localhost:8545'))
    role_pause = w3.keccak(text='ROLE_PAUSE')  # index = 0
    role_mint = w3.keccak(text='ROLE_MINT')  # index = 1
    role_stake = w3.keccak(text='ROLE_STAKE')  # index = 2
    role_fee_manager = w3.keccak(text='ROLE_FEE_MANAGER')  # index = 3
    role_oracle = w3.keccak(text='ROLE_ORACLE')  # index = 4
    role_default_admin = w3.toBytes(hexstr="0x00")  # index = 5
    return [role_pause, role_mint, role_stake, role_fee_manager, role_oracle, role_default_admin]

# Predefined Accounts
@pytest.fixture
def owner():
    return accounts[0]

@pytest.fixture
def deployer():
    return accounts[1]

@pytest.fixture
def admin():
    return accounts[2]

@pytest.fixture
def oracle():
    return accounts[3]

@pytest.fixture
def delegates():
    return [accounts[4], accounts[5]]

@pytest.fixture
def users():
    return [accounts[6], accounts[7]]

# Protocol params
@pytest.fixture
def start_amount():
    return 100

@pytest.fixture
def common_ratio():
    return 10

@pytest.fixture
def sequence_length():
    return 3

@pytest.fixture
def stake_duration():
    return 1

@pytest.fixture
def stake_amounts(start_amount, common_ratio):
    return [start_amount, start_amount * common_ratio, start_amount * common_ratio * common_ratio]

@pytest.fixture
def manager_fee_shares():
    return 100

