package keeper

import (
	"fmt"
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"time"
)

// BeginBlocker of epochs module
func (k Keeper) BeginBlocker(ctx sdk.Context, sk types.StakingKeeper) {

	isWhitelistEnabled := k.GetParams(ctx).WhitelistEnabled
	whitelistCount := k.GetWhiteListedValidatorCount(sdk.WrapSDKContext(ctx))

	//logger := k.Logger(ctx)

	fmt.Println("\n=================================================")
	fmt.Println("PERMISSIONS MODULE - BeginBlocker", time.Now().Format(time.RFC822))
	fmt.Println("is_whitelist_enabled: ", isWhitelistEnabled)
	fmt.Println("whitelistCount: ", whitelistCount)
	fmt.Println("=================================================")

	if isWhitelistEnabled {
		if whitelistCount == 0 {

			fmt.Println("\n=================================================")
			fmt.Println("PERMISSIONS MODULE - BeginBlocker - Append to Whitelist", time.Now().Format(time.RFC822))
			fmt.Println("=================================================")

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

			fmt.Println("\n=================================================")
			fmt.Println("PERMISSIONS MODULE - BeginBlocker - Delete Whitelist", time.Now().Format(time.RFC822))

			delStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
			store := ctx.KVStore(k.storeKey)
			iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhiteListedValidatorKey))
			defer iterator.Close()
			for ; iterator.Valid(); iterator.Next() {
				validator := MustUnmarshalValidator(k.cdc, iterator.Value())
				operatorAddr, _ := sdk.ValAddressFromBech32(validator.ValidatorAddress)
				fmt.Println("operatorAddr: ", operatorAddr)
				delStore.Delete(types.GetValidatorKey(operatorAddr))
			}

			fmt.Println("=================================================")

		}
	}

}
