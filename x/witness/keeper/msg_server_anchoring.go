package keeper

import (
	"context"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func (k msgServer) Anchoring(goCtx context.Context, msg *types.MsgAnchoring) (*types.MsgAnchoringResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var anchored = types.Anchored{
		BlockHash:   msg.BlockHash,
		BlockHeight: msg.BlockHeight,
		ChainID: msg.ChainID,
		Filter: msg.Filter,
	}

	_ = k.AppendAnchored(ctx, anchored)

	return &types.MsgAnchoringResponse{}, nil
}
