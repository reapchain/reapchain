package permissions

import (
	"fmt"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
	"github.com/reapchain/reapchain/v8/x/permissions/keeper"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// NewPermissionsModuleProposalHandler creates a governance handler to manage new proposal types.
func NewPermissionsModuleProposalHandler(k *keeper.Keeper, sk types.StakingKeeper, bk types.BankKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.MsgRegisterStandingMemberProposal:
			return handleRegisterNewStandingMemberProposal(ctx, k, sk, c)
		case *types.MsgRemoveStandingMemberProposal:
			return handleRemoveStandingMemberProposal(ctx, k, sk, bk, c)
		case *types.MsgReplaceStandingMemberProposal:
			return handleReplaceStandingMemberProposal(ctx, k, sk, bk, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterNewStandingMemberProposal(ctx sdk.Context, k *keeper.Keeper, sk types.StakingKeeper, p *types.MsgRegisterStandingMemberProposal) error {

	result := k.ValidateRegisterStandingMember(ctx, sk, p.ValidatorAddress, p.AccountAddress)

	if result != nil {
		return result
	}
	var whiteListedValidator = types.WhitelistedValidator{
		ValidatorAddress: p.ValidatorAddress,
		Moniker:          p.Moniker,
		AccountAddress:   p.AccountAddress,
	}

	err := k.AppendWhitelistedValidator(ctx, whiteListedValidator)
	if err != nil {
		return err
	}

	k.Logger(ctx).Info(
		fmt.Sprintf("attempt to register new validator; validator address: %s, moniker: %s", whiteListedValidator.ValidatorAddress, whiteListedValidator.Moniker),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeMsgRegisterStandingMemberProposal,
			sdk.NewAttribute(types.WhiteListedValidatorKey, whiteListedValidator.ValidatorAddress),
		))
	return nil
}

func handleRemoveStandingMemberProposal(ctx sdk.Context, k *keeper.Keeper, sk types.StakingKeeper, bk types.BankKeeper, p *types.MsgRemoveStandingMemberProposal) error {

	validateResult := k.ValidateRemoveStandingMember(ctx, sk, p.ValidatorAddress)

	if validateResult != nil {
		return validateResult
	}

	forcedUnbondingResult := k.ForceUnbondAllDelegations(ctx, sk, bk, p.ValidatorAddress)
	if forcedUnbondingResult != nil {
		return forcedUnbondingResult
	}

	whiteListedValidator, getValResult := k.GetValidatorWithAddress(ctx, p.ValidatorAddress)

	if getValResult != nil {
		return getValResult
	}

	removeResult := k.RemoveWhitelistedValidator(ctx, whiteListedValidator)
	if removeResult != nil {
		return removeResult
	}

	k.Logger(ctx).Info(
		fmt.Sprintf("attempt to remove validator; validator address: %s, moniker: %s", whiteListedValidator.ValidatorAddress, whiteListedValidator.Moniker),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeMsgRemoveStandingMemberProposal,
			sdk.NewAttribute(types.WhiteListedValidatorKey, whiteListedValidator.ValidatorAddress),
		))
	return nil
}

func handleReplaceStandingMemberProposal(ctx sdk.Context, k *keeper.Keeper, sk types.StakingKeeper, bk types.BankKeeper, p *types.MsgReplaceStandingMemberProposal) error {

	validateResult := k.ValidateRemoveStandingMember(ctx, sk, p.ExistingValidatorAddress)

	if validateResult != nil {
		return validateResult
	}

	forcedUnbondingResult := k.ForceUnbondAllDelegations(ctx, sk, bk, p.ExistingValidatorAddress)
	if forcedUnbondingResult != nil {
		return forcedUnbondingResult
	}

	whiteListedValidator, getValResult := k.GetValidatorWithAddress(ctx, p.ExistingValidatorAddress)

	if getValResult != nil {
		return getValResult
	}

	removeResult := k.RemoveWhitelistedValidator(ctx, whiteListedValidator)
	if removeResult != nil {
		return removeResult
	}

	result := k.ValidateRegisterStandingMember(ctx, sk, p.ReplacementValidatorAddress, p.ReplacementAccountAddress)

	if result != nil {
		return result
	}
	var ReplacementWhiteListedValidator = types.WhitelistedValidator{
		ValidatorAddress: p.ReplacementValidatorAddress,
		Moniker:          p.ReplacementAccountAddress,
		AccountAddress:   p.ReplacementMoniker,
	}

	err := k.AppendWhitelistedValidator(ctx, ReplacementWhiteListedValidator)
	if err != nil {
		return err
	}

	k.Logger(ctx).Info(
		fmt.Sprintf("attempt to repace a validator; existing validator: %s, replacement validator: %s", p.ExistingValidatorAddress, whiteListedValidator.ValidatorAddress),
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeMsgRemoveStandingMemberProposal,
			sdk.NewAttribute(types.WhiteListedValidatorKey, whiteListedValidator.ValidatorAddress),
		))
	return nil
}
