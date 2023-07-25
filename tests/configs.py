import pytest
from brownie import accounts

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

