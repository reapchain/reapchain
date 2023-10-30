package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/crypto/tmhash"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func (k Keeper) GetEscrowPoolByDenom(ctx sdk.Context, denom string) (types.EscrowPool, bool) {
	var escrowSupply types.EscrowPool
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	bz := store.Get(tmhash.Sum([]byte(denom)))
	if len(bz) == 0 {
		return types.EscrowPool{}, false
	}
	k.cdc.MustUnmarshal(bz, &escrowSupply)

	return escrowSupply, true
}

func (k Keeper) GetTotalEscrowPool(ctx sdk.Context) []types.EscrowPool {
	var totalEscrowSupply []types.EscrowPool

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixEscrowSupply)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var escrowPool types.EscrowPool
		k.cdc.MustUnmarshal(iterator.Value(), &escrowPool)
		totalEscrowSupply = append(totalEscrowSupply, escrowPool)
	}

	return totalEscrowSupply
}

func (k Keeper) AddToEscrowPool(ctx sdk.Context, escrowpool types.EscrowPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	key := escrowpool.GetID()
	bz := k.cdc.MustMarshal(&escrowpool)
	store.Set(key, bz)
}

func (k Keeper) SetEscrowPool(ctx sdk.Context, escrowSupply types.EscrowPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	key := escrowSupply.GetID()
	bz := k.cdc.MustMarshal(&escrowSupply)
	store.Set(key, bz)
}
