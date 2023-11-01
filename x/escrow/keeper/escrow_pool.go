package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/crypto/tmhash"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func (k Keeper) GetEscrowPoolByDenom(ctx sdk.Context, denom string) (sdk.Coin, bool) {
	var escrowPool sdk.Coin
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowPool)
	bz := store.Get(tmhash.Sum([]byte(denom)))
	if len(bz) == 0 {
		return sdk.Coin{}, false
	}
	k.cdc.MustUnmarshal(bz, &escrowPool)

	return escrowPool, true
}

func (k Keeper) GetAllEscrowPools(ctx sdk.Context) []sdk.Coin {
	var allEscrowPools []sdk.Coin

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixEscrowPool)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var escrowPool sdk.Coin
		k.cdc.MustUnmarshal(iterator.Value(), &escrowPool)
		allEscrowPools = append(allEscrowPools, escrowPool)
	}

	return allEscrowPools
}

func (k Keeper) SetEscrowPool(ctx sdk.Context, escrowPool sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowPool)
	key := types.GetID(escrowPool)
	bz := k.cdc.MustMarshal(&escrowPool)
	store.Set(key, bz)
}
