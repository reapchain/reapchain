package keeper

import (
	"fmt"

	"github.com/reapchain/reapchain-core/libs/log"

	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/bech32ibc/types"
)

type (
	Keeper struct {
		channelKeeper types.ChannelKeeper

		cdc      codec.Codec
		storeKey sdk.StoreKey

		tk types.TransferKeeper
	}
)

func NewKeeper(
	channelKeeper types.ChannelKeeper,
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	tk types.TransferKeeper,
) *Keeper {
	return &Keeper{
		channelKeeper: channelKeeper,
		cdc:           cdc,
		storeKey:      storeKey,
		tk:            tk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
