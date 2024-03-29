package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/reapchain/v8/x/recovery/types"
)

// GetParams returns the total set of recovery parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSetIfExists(ctx, &params)
	return params
}

// SetParams sets the recovery parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
