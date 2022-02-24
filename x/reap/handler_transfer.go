package reap

import (
	"github.com/reapchain/mercury/x/reap/keeper"
	"github.com/reapchain/mercury/x/reap/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
)

func handleTransferMintedCoins(ctx sdk.Context, k keeper.Keeper, msg *types.TransferMintedCoins) (*sdk.Result, error) {
	k.TransferMintedCoins(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}