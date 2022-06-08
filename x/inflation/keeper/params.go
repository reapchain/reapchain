package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/tharsis/evmos/v4/x/inflation/types"
)

// GetParams returns the total set of inflation parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the inflation parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
