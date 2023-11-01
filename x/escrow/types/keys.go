package types

import (
	"github.com/ethereum/go-ethereum/common"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"
)

// constants
const (
	// module name
	ModuleName = "escrow"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName

	DenomPrefix = "gravity"
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

// prefix bytes for the EVM persistent store
const (
	prefixDenom = iota + 1
	prefixEscrowPool
)

// KVStore key prefixes
var (
	KeyPrefixDenom      = []byte{prefixDenom}
	KeyPrefixEscrowPool = []byte{prefixEscrowPool}
	KeyPrefixTokenPair  = []byte{prefixDenom}
)
