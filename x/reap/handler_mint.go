package reap

import (
	"github.com/reapchain/reapchain/x/reap/keeper"
	"github.com/reapchain/reapchain/x/reap/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
)

func handleMintCoins(ctx sdk.Context, k keeper.Keeper, msg *types.MintCoins) (*sdk.Result, error) {
	k.MintCoins(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}