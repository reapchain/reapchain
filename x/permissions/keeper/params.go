package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.WhitelistEnabled(ctx),
		k.GovMinInitialDepositEnabled(ctx),
		k.GovMinInitialDepositPercentage(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) WhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyPodcWhitelistEnabled, &res)
	return
}

// GovMinInitialDepositEnabled  returns the GovernanceMinimumInitialDepositEnabled param
func (k Keeper) GovMinInitialDepositEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyGovMinInitialDepositEnabled, &res)
	return
}

// GovMinInitialDepositPercentage returns the GovernanceMinimumInitialDepositPercentage param
func (k Keeper) GovMinInitialDepositPercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyGovMinInitialDepositPercentage, &res)
	return
}

// GetIfExistsWhitelistEnabled returns the WhitelistEnabled param
func (k Keeper) GetIfExistsWhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.GetIfExists(ctx, types.KeyPodcWhitelistEnabled, &res)
	return
}
