package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v4/x/inflation/types"
)

// GetMaxCoins gets current maxCoins
func (k Keeper) GetMaxCoins(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixMaxCoins)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}

// SetMaxCoins stores the current maxCoins
func (k Keeper) SetMaxCoins(ctx sdk.Context, maxCoins string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixMaxCoins, []byte(maxCoins))
}

// GetCurrKeyPrefixCurrentInflation gets current inflations
func (k Keeper) GetCurrentInflation(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixCurrentInflation)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}

// SetCurrKeyPrefixCurrentInflation stores the current inflations
func (k Keeper) SetCurrentInflation(ctx sdk.Context, inflations string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixCurrentInflation, []byte(inflations))
}