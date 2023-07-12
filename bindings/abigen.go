//go:generate abigen --abi ../build/contracts/SystemStaking.abi --pkg bindings --out SystemStaking.go --type SystemStaking
//go:generate abigen --abi ../build/contracts/UniIOTX.abi --pkg bindings --out UniIOTX.go --type UniIOTX
//go:generate abigen --abi ../build/contracts/IOTXClear.abi --pkg bindings --out IOTXClear.go --type IOTXClear
//go:generate abigen --abi ../build/contracts/IOTXStake.abi --pkg bindings --out IOTXStake.go --type IOTXStake

package bindings

// Compatible abigen version: 1.10.17-stable. Source code download links:
// https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.10.17.tar.gz
