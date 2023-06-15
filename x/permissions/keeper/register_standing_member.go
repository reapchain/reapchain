package keeper

import (
	"fmt"
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"time"
)

func (k Keeper) AppendWhitelistedValidator(
	ctx sdk.Context,
	whiteListedValidator types.WhitelistedValidator,
) error {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
	appendedValue := k.cdc.MustMarshal(&whiteListedValidator)
	operatorAddr, err := sdk.ValAddressFromBech32(whiteListedValidator.ValidatorAddress)

	fmt.Println("\n=================================================")
	fmt.Println("PERMISSIONS MODULE - AppendWhitelistedValidator", time.Now().Format(time.RFC822))
	fmt.Println("operatorAddr: ", operatorAddr)
	fmt.Println("whiteListedValidator: ", whiteListedValidator)
	fmt.Println("=================================================")

	if err != nil {
		return err
	}
	store.Set(types.GetValidatorKey(operatorAddr), appendedValue)
	return nil
}

func (k Keeper) ValidateRegisterStandingMember(ctx sdk.Context, sk types.StakingKeeper, valAddrStr string, accAddrStr string) error {

	maxStandingMembers := sk.MaxStandingMembers(ctx)
	currentWhitelistCount := k.GetWhiteListedValidatorCount(sdk.WrapSDKContext(ctx))

	fmt.Println("\n==================================================")
	fmt.Println("PERMISSIONS MODULE - ValidateRegisterStandingMember", time.Now().Format(time.RFC822))
	fmt.Println("max_standing_members: ", maxStandingMembers)
	fmt.Println("current_whitelist_count: ", currentWhitelistCount)

	if currentWhitelistCount == maxStandingMembers {
		return types.ErrMaximumNumberOfStandingMember
	}

	validatorAddress, _ := sdk.ValAddressFromBech32(valAddrStr)
	accAddressFromValidatorAddress := sdk.AccAddress(validatorAddress)

	fmt.Println("validatorAddress: ", validatorAddress.String())

	fmt.Println("accAddressFromValidatorAddress: ", accAddressFromValidatorAddress.String())
	fmt.Println("accAddrStr: ", accAddrStr)

	if accAddressFromValidatorAddress.String() != accAddrStr {
		return types.ErrAccountsDontMatch
	}

	found := k.FindValidator(ctx, validatorAddress)

	if found {
		return types.ErrValidatorAlreadyRegistered
	}

	fmt.Println("==================================================")

	return nil
}
