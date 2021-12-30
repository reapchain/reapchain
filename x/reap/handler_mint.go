package reap

import (
	"github.com/tharsis/evmos/x/reap/keeper"
	"github.com/tharsis/evmos/x/reap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMintCoins(ctx sdk.Context, k keeper.Keeper, msg *types.MintCoins) (*sdk.Result, error) {
	k.MintCoins(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}