package keeper

import (
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

func (k Keeper) AppendWhitelistedValidator(
	ctx sdk.Context,
	whiteListedValidator types.WhitelistedValidator,
) error {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
	appendedValue := k.cdc.MustMarshal(&whiteListedValidator)
	operatorAddr, err := sdk.ValAddressFromBech32(whiteListedValidator.ValidatorAddress)

	if err != nil {
		return err
	}
	store.Set(types.GetValidatorKey(operatorAddr), appendedValue)
	return nil
}

func (k Keeper) ValidateRegisterStandingMember(ctx sdk.Context, sk types.StakingKeeper, valAddrStr string, accAddrStr string) error {

	maxStandingMembers := sk.MaxStandingMembers(ctx)
	currentWhitelistCount := k.GetWhiteListedValidatorCount(sdk.WrapSDKContext(ctx))

	if currentWhitelistCount == maxStandingMembers {
		return types.ErrMaximumNumberOfStandingMember
	}

	validatorAddress, _ := sdk.ValAddressFromBech32(valAddrStr)
	accAddressFromValidatorAddress := sdk.AccAddress(validatorAddress)

	if accAddressFromValidatorAddress.String() != accAddrStr {
		return types.ErrAccountsDontMatch
	}

	found := k.FindValidator(ctx, validatorAddress)

	if found {
		return types.ErrValidatorAlreadyRegistered
	}

	return nil
}
