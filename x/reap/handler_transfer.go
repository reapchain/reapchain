package reap

import (
	"github.com/tharsis/evmos/x/reap/keeper"
	"github.com/tharsis/evmos/x/reap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleTransferMintedCoins(ctx sdk.Context, k keeper.Keeper, msg *types.TransferMintedCoins) (*sdk.Result, error) {
	k.TransferMintedCoins(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}