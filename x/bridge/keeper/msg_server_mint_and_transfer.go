package keeper

import (
	"context"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/bridge/types"
)

func (k msgServer) MintAndTransfer(goCtx context.Context, msg *types.MsgMintAndTransfer) (*types.MsgMintAndTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	coins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return &types.MsgMintAndTransferResponse{}, err
	}

	mintError := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if mintError != nil {
		return &types.MsgMintAndTransferResponse{}, mintError
	}
	recipientAddress, addressParseError := sdk.AccAddressFromBech32(msg.Address)
	if addressParseError != nil {
		return &types.MsgMintAndTransferResponse{}, addressParseError
	}

	transferError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddress, coins)
	if transferError != nil {
		return &types.MsgMintAndTransferResponse{}, transferError
	}

	return &types.MsgMintAndTransferResponse{}, nil
}
