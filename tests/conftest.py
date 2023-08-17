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

@pytest.fixture(scope="session", autouse=True)
def uint256_max():
    return 115792089237316195423570985008687907853269984665640564039457584007913129639935

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
    return accounts.at("0x065e1164818487818E6BA714E8d80B91718ad758", True)

@pytest.fixture(scope="session", autouse=True)
def deployer():
    return accounts.at("0x3af43AfEd31C00311381A8DF26cc58C9bD2b5375", True)

@pytest.fixture(scope="session", autouse=True)
def admin():
    return accounts.at("0xbFdDf5e269C74157b157c7DaC5E416d44afB790d", True)

@pytest.fixture(scope="session", autouse=True)
def oracles():
    oracle0 = accounts.at("0xC8a85eD8A9aBF0a21031B7c62C13464D1527cd09", True)
    oracle1 = accounts.at("0x912AD2282799C5d62334017578418471f5aF5353", True)
    return [oracle0, oracle1]

@pytest.fixture(scope="session", autouse=True)
def delegates():
    delegate0 = accounts.at("0xac82586b613d10a33df00835aC6DAd8541952334", True)
    delegate1 = accounts.at("0xE88eBFccF58aaf553134AE5f87a77d0608B76d53", True)
    return [delegate0, delegate1]

@pytest.fixture(scope="session", autouse=True)
def users():
    user0 = accounts.at("0x9ACE9968545089893208f35A81569Fa81cd24F7c", True)
    user1 = accounts.at("0x912AD2282799C5d62334017578418471f5aF5353", True)
    return [user0, user1]

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

@pytest.fixture
def fund(owner, deployer, admin, oracles, delegates, users):
    extra_accounts = [owner, deployer, admin, oracles[0], oracles[1], delegates[0], delegates[1], users[0], users[1]]
    for account in extra_accounts:
        accounts[0].transfer(account, "100000000 ether")
        print("Balance:", account.balance())


# Contracts
@pytest.fixture()
def contracts(proxy, fund, owner, deployer, admin, oracles, delegates, stake_duration, start_amount, common_ratio, sequence_length, stake_amounts, manager_fee_shares):
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
        oracles[0],
        start_amount,
        common_ratio,
        sequence_length,
        stake_duration,
        {'from': admin}
    )
    iotx_clear_transparent.initialize(system_staking, iotx_staking_transparent, oracles[0], {'from': admin})

    iotx_staking_transparent.setManagerFeeShares(manager_fee_shares, {'from': admin})
    iotx_staking_transparent.setGlobalDelegate(delegates[0], {'from': oracles[0]})

    return [system_staking, uni_iotx_transparent, iotx_clear_transparent, iotx_staking_transparent]
