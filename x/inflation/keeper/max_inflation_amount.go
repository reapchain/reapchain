package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/inflation/types"
)

// GetMaxInflationAmount gets current maxInflationAmount
func (k Keeper) GetMaxInflationAmount(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.ParamStoreKeyMaxInflationAmount, &res)
	return
}

// SetMaxInflationAmount stores the current maxInflationAmount
func (k Keeper) SetMaxInflationAmount(ctx sdk.Context, maxInflationAmount string) {
	if !k.paramstore.HasKeyTable() {
		ps := k.paramstore.WithKeyTable(types.ParamKeyTable())
		k.paramstore = ps
	}

	k.paramstore.Set(ctx, types.ParamStoreKeyMaxInflationAmount, maxInflationAmount)
}

// GetCurrKeyPrefixCurrentInflation gets current inflations
func (k Keeper) GetCurrentInflationAmount(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.ParamStoreKeyCurrentInflationAmount, &res)
	return
}

// SetCurrKeyPrefixCurrentInflation stores the current inflations
func (k Keeper) SetCurrentInflation(ctx sdk.Context, currentInflationsAmount string) {
	if !k.paramstore.HasKeyTable() {
		ps := k.paramstore.WithKeyTable(types.ParamKeyTable())
		k.paramstore = ps
	}

	k.paramstore.Set(ctx, types.ParamStoreKeyCurrentInflationAmount, currentInflationsAmount)
}
