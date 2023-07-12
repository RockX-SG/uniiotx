from brownie import SystemStaking, UniIOTX, IOTXClear, IOTXStake, accounts

# DefaultAdmin address: "0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC"
# DefaultAdmin private key: "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027"

def main():
    account = accounts.load('DefaultAdmin')
    SystemStaking.deploy({'from': account})
    UniIOTX.deploy({'from': account})
    IOTXClear.deploy({'from': account})
    IOTXStake.deploy({'from': account})
