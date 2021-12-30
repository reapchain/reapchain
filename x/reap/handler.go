package reap

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/evmos/x/reap/keeper"
	"github.com/tharsis/evmos/x/reap/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MintCoins:
			return handleMintCoins(ctx, k, msg)

		case *types.BurnCoins:
			return handleBurnCoins(ctx, k, msg)

		case *types.TransferMintedCoins:
			return handleTransferMintedCoins(ctx, k, msg)

		case *types.BurnCoinsFromAccount:
			return handleBurnCoinsFromAccount(ctx, k, msg)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
