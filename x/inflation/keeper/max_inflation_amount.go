package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/inflation/types"
)

// GetMaxInflationAmount gets current maxInflationAmount
func (k Keeper) GetMaxInflationAmount(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.MaxInflationAmount
}

// SetMaxInflationAmount stores the current maxInflationAmount
func (k Keeper) SetMaxInflationAmount(ctx sdk.Context, maxInflationAmount string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ParamStoreKeyMaxInflationAmount, []byte(maxInflationAmount))
}

// GetCurrKeyPrefixCurrentInflation gets current inflations
func (k Keeper) GetCurrentInflationAmount(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixCurrentInflationAmount)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}

// SetCurrKeyPrefixCurrentInflation stores the current inflations
func (k Keeper) SetCurrentInflation(ctx sdk.Context, inflations string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixCurrentInflationAmount, []byte(inflations))
}
