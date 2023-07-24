package keeper

import (
	"time"

	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

func (k Keeper) RemoveWhitelistedValidator(
	ctx sdk.Context,
	whiteListedValidator types.WhitelistedValidator,
) error {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
	operatorAddr, err := sdk.ValAddressFromBech32(whiteListedValidator.ValidatorAddress)
	if err != nil {
		return err
	}
	store.Delete(types.GetValidatorKey(operatorAddr))

	return nil
}

func (k Keeper) ValidateRemoveStandingMember(ctx sdk.Context, sk types.StakingKeeper, valAddrStr string) error {

	validatorAddress, err := sdk.ValAddressFromBech32(valAddrStr)
	if err != nil {
		return types.ErrInvalidValidatorAddress
	}
	found := k.FindValidator(ctx, validatorAddress)
	if !found {
		return types.ErrValidatorNotFound
	}

	return nil
}

func (k Keeper) ForceUnbondAllDelegations(ctx sdk.Context, sk types.StakingKeeper, bk types.BankKeeper, valAddrStr string) error {

	validatorAddress, _ := sdk.ValAddressFromBech32(valAddrStr)

	validator := sk.Validator(ctx, validatorAddress)

	if validator.IsJailed() {
		consAddr, err := validator.GetConsAddr()
		if err != nil {
			return err
		}
		sk.Unjail(ctx, consAddr)
	}

	listOfDelegations := sk.GetValidatorDelegations(ctx, validatorAddress)
	for _, delegation := range listOfDelegations {
		delegatorShares := delegation.GetShares()
		delegatorAddress := delegation.GetDelegatorAddr()

		validator, found := sk.GetValidator(ctx, validatorAddress)
		if !found {
			return stakingtypes.ErrNoValidatorFound
		}
		if sk.HasMaxUnbondingDelegationEntries(ctx, delegatorAddress, validatorAddress) {
			return stakingtypes.ErrMaxUnbondingDelegationEntries
		}

		returnAmount, err := sk.Unbond(ctx, delegatorAddress, validatorAddress, delegatorShares)
		if err != nil {
			return err
		}

		if validator.IsBonded() {
			coins := sdk.NewCoins(sdk.NewCoin(sk.BondDenom(ctx), returnAmount))
			if err := bk.SendCoinsFromModuleToModule(ctx, stakingtypes.BondedPoolName, stakingtypes.NotBondedPoolName, coins); err != nil {
				panic(err)
			}
		}

		unbondingDuration, _ := time.ParseDuration(types.DefaultUnbondingTime)
		completionTime := ctx.BlockHeader().Time.Add(unbondingDuration)

		ubd := sk.SetUnbondingDelegationEntry(ctx, delegatorAddress, validatorAddress, ctx.BlockHeight(), completionTime, returnAmount)
		sk.InsertUBDQueue(ctx, ubd, completionTime)

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				stakingtypes.EventTypeUnbond,
				sdk.NewAttribute(stakingtypes.AttributeKeyValidator, delegatorAddress.String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, delegatorShares.String()),
				sdk.NewAttribute(stakingtypes.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),
			),
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, stakingtypes.AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeySender, delegatorAddress.String()),
			),
		})
	}
	return nil
}
