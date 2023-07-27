package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// BeginBlocker of permissions module
func (k Keeper) BeginBlocker(ctx sdk.Context, sk types.StakingKeeper) {

	isWhitelistEnabled := k.GetParams(ctx).PodcWhitelistEnabled
	whitelistCount := k.GetWhiteListedValidatorCount(sdk.WrapSDKContext(ctx))
	standingMemberCount := sk.CountStandingMember(ctx)

	if isWhitelistEnabled {

		validators := sk.GetAllValidators(ctx)

		if whitelistCount == 0 {
			for _, validator := range validators {
				if validator.Type == stakingtypes.ValidatorTypeStanding {

					if !validator.DelegatorShares.IsZero() && !validator.IsUnbonding() {
						validatorAddress, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
						accountAddress := sdk.AccAddress(validatorAddress).String()
						var whiteListedValidator = types.WhitelistedValidator{
							ValidatorAddress: validator.OperatorAddress,
							Moniker:          validator.GetMoniker(),
							AccountAddress:   accountAddress,
						}
						k.AppendWhitelistedValidator(ctx, whiteListedValidator)
					}
				}
			}
		}

		if standingMemberCount > whitelistCount {

			for _, validator := range validators {

				if validator.Type == stakingtypes.ValidatorTypeStanding {

					validatorAddress, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
					foundInWhitelist := k.FindValidator(ctx, validatorAddress)

					if !foundInWhitelist {

						blockHeight := ctx.BlockHeight()
						unbondingTime := validator.UnbondingTime
						remainingUnbondingTime := unbondingTime.Sub(ctx.BlockTime())
						var unbondingStartHeight = validator.UnbondingHeight

						if validator.DelegatorShares.IsZero() &&
							validator.IsUnbonding() &&
							remainingUnbondingTime > types.DefaultSafeTimeInterval &&
							blockHeight > (unbondingStartHeight+types.DefaultRemovalBlockInterval) {

							validator = sk.UnbondingToUnbonded(ctx, validator)
							sk.RemoveValidator(ctx, validator.GetOperator())
							sk.DeleteValidatorQueue(ctx, validator)
						}
					}
				}
			}
		}

	} else {
		if whitelistCount > 0 {

			delStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
			store := ctx.KVStore(k.storeKey)
			iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhiteListedValidatorKey))
			defer iterator.Close()
			for ; iterator.Valid(); iterator.Next() {
				validator := MustUnmarshalValidator(k.cdc, iterator.Value())
				operatorAddr, _ := sdk.ValAddressFromBech32(validator.ValidatorAddress)
				delStore.Delete(types.GetValidatorKey(operatorAddr))
			}
		}
	}
}
