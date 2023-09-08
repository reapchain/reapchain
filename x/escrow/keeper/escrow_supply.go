package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/crypto/tmhash"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func (k Keeper) GetEscrowSupplyByDenom(ctx sdk.Context, denom string) (types.EscrowSupply, bool) {
	var escrowSupply types.EscrowSupply
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	bz := store.Get(tmhash.Sum([]byte(denom)))
	if len(bz) == 0 {
		return types.EscrowSupply{}, false
	}
	k.cdc.MustUnmarshal(bz, &escrowSupply)

	return escrowSupply, true
}

func (k Keeper) GetTotalEscrowSupply(ctx sdk.Context) []types.EscrowSupply {
	var totalEscrowSupply []types.EscrowSupply

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixEscrowSupply)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var escrowSupply types.EscrowSupply
		k.cdc.MustUnmarshal(iterator.Value(), &escrowSupply)
		totalEscrowSupply = append(totalEscrowSupply, escrowSupply)
	}

	return totalEscrowSupply
}

func (k Keeper) AddEscrowSupply(ctx sdk.Context, escrowsupply types.EscrowSupply) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	key := escrowsupply.GetID()
	bz := k.cdc.MustMarshal(&escrowsupply)
	store.Set(key, bz)
}

func (k Keeper) SetEscrowSupply(ctx sdk.Context, escrowSupply types.EscrowSupply) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEscrowSupply)
	key := escrowSupply.GetID()
	bz := k.cdc.MustMarshal(&escrowSupply)
	store.Set(key, bz)
}
