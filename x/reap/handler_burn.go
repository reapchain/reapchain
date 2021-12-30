package reap

import (
	"github.com/tharsis/evmos/x/reap/keeper"
	"github.com/tharsis/evmos/x/reap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleBurnCoins(ctx sdk.Context, k keeper.Keeper, msg *types.BurnCoins) (*sdk.Result, error) {
	k.BurnCoins(ctx, *msg)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleBurnCoinsFromAccount(ctx sdk.Context, k keeper.Keeper, msg *types.BurnCoinsFromAccount) (*sdk.Result, error) {
	k.BurnCoinsFromAccount(ctx, *msg)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}