package types

import (
	"math/big"

	sdk "github.com/reapchain/cosmos-sdk/types"
)

const (
	// AttoReap defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	AttoReap string = "areap"

	// BaseDenomUnit defines the base denomination unit for Reaps.
	// 1 reap = 1x10^{BaseDenomUnit} areap
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewReapCoin is a utility function that returns an "areap" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewReapCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(AttoReap, amount)
}

// NewReapDecCoin is a utility function that returns an "areap" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewReapDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoReap, amount)
}

// NewReapCoinInt64 is a utility function that returns an "areap" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewReapCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoReap, amount)
}
