package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.WhitelistEnabled(ctx),
		k.ForcedUnbondingTime(ctx),
		k.GovernanceMinimumInitialDepositEnabled(ctx),
		k.GovernanceMinimumInitialDepositPercentage(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// ForcedUnbondingTime returns the ForcedUnbondingTime param
func (k Keeper) WhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyWhitelistEnabled, &res)
	return
}

// ForcedUnbondingTime returns the ForcedUnbondingTime param
func (k Keeper) ForcedUnbondingTime(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyForcedUnbondingTime, &res)
	return
}

// GovernanceMinimumInitialDepositEnabled returns the GovernanceMinimumInitialDepositEnabled param
func (k Keeper) GovernanceMinimumInitialDepositEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyGovernanceMinimumInitialDepositEnabled, &res)
	return
}

// GovernanceMinimumInitialDepositPercentage returns the GovernanceMinimumInitialDepositPercentage param
func (k Keeper) GovernanceMinimumInitialDepositPercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyGovernanceMinimumInitialDepositPercentage, &res)
	return
}

// ForcedUnbondingTime returns the ForcedUnbondingTime param
func (k Keeper) GetIfExistsWhitelistEnabled(ctx sdk.Context) (res bool) {
	k.paramstore.GetIfExists(ctx, types.KeyWhitelistEnabled, &res)
	return
}
