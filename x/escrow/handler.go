package escrow

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// NewHandler defines the escrow module handler instance
func NewHandler(server types.MsgServer) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgConvertToNative:
			res, err := server.ConvertToNative(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgConvertToDenom:
			res, err := server.ConvertToDenom(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			err := sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, err
		}
	}
}
