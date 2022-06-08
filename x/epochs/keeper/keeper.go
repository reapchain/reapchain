package keeper

import (
	"fmt"

	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/libs/log"

	"github.com/tharsis/evmos/v4/x/epochs/types"
)

// Keeper of this module maintains collections of epochs and hooks.
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
	hooks    types.EpochHooks
}

// NewKeeper returns a new instance of epochs Keeper
func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// SetHooks set the epoch hooks
func (k *Keeper) SetHooks(eh types.EpochHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set epochs hooks twice")
	}

	k.hooks = eh

	return k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
