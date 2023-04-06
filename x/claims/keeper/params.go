package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/reapchain/v8/x/claims/types"
)

// GetParams returns the total set of claim parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSetIfExists(ctx, &params)
	return params
}

// SetParams sets the claim parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
