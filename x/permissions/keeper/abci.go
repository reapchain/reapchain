package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// BeginBlocker of permissions module
func (k Keeper) BeginBlocker(ctx sdk.Context, sk types.StakingKeeper) {

	isWhitelistEnabled := k.GetParams(ctx).WhitelistEnabled
	whitelistCount := k.GetWhiteListedValidatorCount(sdk.WrapSDKContext(ctx))

	if isWhitelistEnabled {
		if whitelistCount == 0 {

			validators := sk.GetAllValidators(ctx)
			for _, validator := range validators {

				if validator.Type == "standing" {
					validatorAddress, _ := sdk.ValAddressFromBech32(validator.OperatorAddress)
					accountAddress := sdk.AccAddress(validatorAddress).String()

					var whiteListedValidator = types.WhitelistedValidator{
						ValidatorAddress: validator.OperatorAddress,
						Moniker:          validator.GetMoniker(),
						AccountAddress:   accountAddress,
					}
					_ = k.AppendWhitelistedValidator(
						ctx,
						whiteListedValidator,
					)
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
