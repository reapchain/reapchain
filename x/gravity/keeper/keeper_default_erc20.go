package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/gravity/types"
)

// Set DefaultErc20ContractAddress
func (k Keeper) SetDefaultErc20ContractAddress(ctx sdk.Context, v string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DefaultErc20ContractAddressKey, []byte(v))
}

// Get DefaultErc20ContractAddress
func (k Keeper) GetDefaultErc20ContractAddress(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.DefaultErc20ContractAddressKey)
	if len(bz) == 0 {
		return ""
	}

	return string(bz)
}
