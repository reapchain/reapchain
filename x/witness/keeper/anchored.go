package keeper

import (
	"encoding/binary"

	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

// GetAnchoredCount get the total number of anchored
func (k Keeper) GetAnchoredCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnchoredCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAnchoredCount set the total number of anchored
func (k Keeper) SetAnchoredCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AnchoredCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAnchored appends a anchored in the store with a new id and update the count
func (k Keeper) AppendAnchored(
	ctx sdk.Context,
	anchored types.Anchored,
) uint64 {
	// Create the anchored
	count := k.GetAnchoredCount(ctx)

	// Set the ID of the appended value
	anchored.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnchoredKey))
	appendedValue := k.cdc.MustMarshal(&anchored)
	store.Set(GetAnchoredIDBytes(anchored.Id), appendedValue)

	// Update anchored count
	k.SetAnchoredCount(ctx, count+1)

	return count
}

// SetAnchored set a specific anchored in the store
func (k Keeper) SetAnchored(ctx sdk.Context, anchored types.Anchored) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnchoredKey))
	b := k.cdc.MustMarshal(&anchored)
	store.Set(GetAnchoredIDBytes(anchored.Id), b)
}

// GetAnchored returns a anchored from its id
func (k Keeper) GetAnchored(ctx sdk.Context, id uint64) (val types.Anchored, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnchoredKey))
	b := store.Get(GetAnchoredIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAnchored removes a anchored from the store
func (k Keeper) RemoveAnchored(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnchoredKey))
	store.Delete(GetAnchoredIDBytes(id))
}

// GetAllAnchored returns all anchored
func (k Keeper) GetAllAnchored(ctx sdk.Context) (list []types.Anchored) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AnchoredKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Anchored
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAnchoredIDBytes returns the byte representation of the ID
func GetAnchoredIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAnchoredIDFromBytes returns ID in uint64 format from a byte array
func GetAnchoredIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
