//go:generate abigen --abi ../abis/SystemStaking.abi --pkg bindings --out SystemStaking.go --type SystemStaking
//go:generate abigen --abi ../abis/UniIOTX.abi --pkg bindings --out UniIOTX.go --type UniIOTX
//go:generate abigen --abi ../abis/IOTXClear.abi --pkg bindings --out IOTXClear.go --type IOTXClear
//go:generate abigen --abi ../abis/IOTXStake.abi --pkg bindings --out IOTXStake.go --type IOTXStake

package bindings

// Compatible abigen version: 1.10.17-stable. Source code download links:
// https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.10.17.tar.gz
