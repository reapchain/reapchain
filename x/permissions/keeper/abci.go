package keeper

import (
	"fmt"
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// BeginBlocker of permissions module
func (k Keeper) BeginBlocker(ctx sdk.Context, sk types.StakingKeeper) {

	isWhitelistEnabled := k.GetParams(ctx).WhitelistEnabled
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
			fmt.Println("\n\n=====================================================")
			fmt.Println("whitelistCount: ", whitelistCount)
			fmt.Println("standingMemberCount: ", standingMemberCount)

			for _, validator := range validators {

				if validator.Type == stakingtypes.ValidatorTypeStanding {

					validatorAddress, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
					foundInWhitelist := k.FindValidator(ctx, validatorAddress)

					bondStatus := validator.GetStatus()
					blockHeight := ctx.BlockHeight()
					unbondingTime := validator.UnbondingTime

					if !foundInWhitelist {
						fmt.Println("\n\n========================     STAKING VALIDATORS     =============================")
						fmt.Println("----     NOT IN WHITELIST     ----")

						fmt.Println("validatorAddress - hex: ", validatorAddress)
						fmt.Println("validatorAddress - string: ", validatorAddress.String())
						fmt.Println("delegation shares: ", validator.DelegatorShares.String())

						fmt.Println("found in whitelist: ", foundInWhitelist)
						fmt.Println("bondStatus: ", bondStatus)
						fmt.Println("blockHeight: ", blockHeight)

						if validator.DelegatorShares.IsZero() && validator.IsUnbonding() {

							if validator.IsUnbonding() {
								var unbondingStartHeight = validator.UnbondingHeight
								fmt.Println("unbonding time: ", unbondingTime)
								fmt.Println("unbondingStartHeight: ", unbondingStartHeight)

								if blockHeight > (unbondingStartHeight + int64(types.DefaultRemovalBlockInterval)) {
									validator = sk.UnbondingToUnbonded(ctx, validator)
									sk.RemoveValidator(ctx, validator.GetOperator())
								}
							}
						}
					}
				}
			}
			fmt.Println("=====================================================\n\n")

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
