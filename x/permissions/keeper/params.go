package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.WhitelistEnabled(ctx),
		k.PermissionsUnbondingTime(ctx),
		k.PermissionsMinimumInitialDepositEnabled(ctx),
		k.PermissionsMinimumInitialDepositPercentage(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) WhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyWhitelistEnabled, &res)
	return
}

// PermissionsUnbondingTime returns the PermissionsUnbondingTime param
func (k Keeper) PermissionsUnbondingTime(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyPermissionsUnbondingTime, &res)
	return
}

// PermissionsMinimumInitialDepositEnabled  returns the GovernanceMinimumInitialDepositEnabled param
func (k Keeper) PermissionsMinimumInitialDepositEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyPermissionsMinimumInitialDepositEnabled, &res)
	return
}

// PermissionsMinimumInitialDepositPercentage returns the GovernanceMinimumInitialDepositPercentage param
func (k Keeper) PermissionsMinimumInitialDepositPercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyPermissionsMinimumInitialDepositPercentage, &res)
	return
}

// GetIfExistsWhitelistEnabled returns the WhitelistEnabled param
func (k Keeper) GetIfExistsWhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.GetIfExists(ctx, types.KeyWhitelistEnabled, &res)
	return
}
