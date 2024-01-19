package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/crypto/tmhash"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func NewEscrowDenom(denom string, enabled bool) types.RegisteredDenom {
	return types.RegisteredDenom{
		Denom:   denom,
		Enabled: true,
	}
}

func (k Keeper) RegisterDenom(ctx sdk.Context, denom types.RegisteredDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDenom)
	key := denom.GetID()
	bz := k.cdc.MustMarshal(&denom)
	store.Set(key, bz)
}

func (k Keeper) GetRegisteredDenoms(ctx sdk.Context) []types.RegisteredDenom {
	var registeredDenoms []types.RegisteredDenom

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixDenom)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var registeredDenom types.RegisteredDenom
		k.cdc.MustUnmarshal(iterator.Value(), &registeredDenom)
		registeredDenoms = append(registeredDenoms, registeredDenom)
	}

	return registeredDenoms
}

func (k Keeper) GetRegisteredDenom(ctx sdk.Context, denom string) (types.RegisteredDenom, bool) {
	if denom == "" {
		return types.RegisteredDenom{}, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDenom)
	var registeredDenom types.RegisteredDenom
	bz := store.Get(tmhash.Sum([]byte(denom)))
	if len(bz) == 0 {
		return types.RegisteredDenom{}, false
	}

	k.cdc.MustUnmarshal(bz, &registeredDenom)
	return registeredDenom, true
}

func (k Keeper) SetRegisteredDenom(ctx sdk.Context, denom types.RegisteredDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	key := denom.GetID()
	bz := k.cdc.MustMarshal(&denom)
	store.Set(key, bz)
}
 
func (k Keeper) IsDenomRegistered(ctx sdk.Context, denom string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixDenom)
	return store.Has(tmhash.Sum([]byte(denom)))
}
