package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/inflation/types"
)

// GetPeriod gets current period
func (k Keeper) GetPeriod(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixPeriod)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetPeriod stores the current period
func (k Keeper) SetPeriod(ctx sdk.Context, period uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixPeriod, sdk.Uint64ToBigEndian(period))
}
