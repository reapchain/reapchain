package ante

import (
	"fmt"
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/x/authz"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	permissionsmodulekeeper "github.com/reapchain/reapchain/v8/x/permissions/keeper"
	permissionstypes "github.com/reapchain/reapchain/v8/x/permissions/types"

	vestingtypes "github.com/reapchain/reapchain/v8/x/vesting/types"
)

// CreateValidatorMessage applies a WhiteList Filtering for Standing Member Creation
type CreateValidatorMessage struct {
	sk  vestingtypes.StakingKeeper
	cdc codec.BinaryCodec
	pk  permissionsmodulekeeper.Keeper
}

// NewCreateValidatorMessage creates a new CreateValidatorMessage
func NewCreateValidatorMessage(sk vestingtypes.StakingKeeper, cdc codec.BinaryCodec, pk permissionsmodulekeeper.Keeper) CreateValidatorMessage {
	return CreateValidatorMessage{
		sk:  sk,
		cdc: cdc,
		pk:  pk,
	}
}

// AnteHandle checks if the tx contains a staking delegation.
// It errors if the coins are still locked or the bond amount is greater than
// the coins already vested
func (sd CreateValidatorMessage) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *authz.MsgExec:
			// Check for bypassing authorization
			if err := sd.validateAuthz(ctx, msg); err != nil {
				return ctx, err
			}
		default:
			if err := sd.validateMsg(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}

// validateAuthz validates the authorization internal message
func (sd CreateValidatorMessage) validateAuthz(ctx sdk.Context, execMsg *authz.MsgExec) error {
	for _, v := range execMsg.Msgs {
		var innerMsg sdk.Msg
		if err := sd.cdc.UnpackAny(v, &innerMsg); err != nil {
			return sdkerrors.Wrap(err, "cannot unmarshal authz exec msgs")
		}
		if err := sd.validateMsg(ctx, innerMsg); err != nil {
			return err
		}
	}

	return nil
}

// validateMsg checks that the only vested coins can be delegated
func (sd CreateValidatorMessage) validateMsg(ctx sdk.Context, msg sdk.Msg) error {

	createValidatorMsg, ok := msg.(*stakingtypes.MsgCreateValidator)
	if !ok {
		return nil
	}

	isWhiteListEnabled := sd.pk.GetIfExistsWhitelistEnabled(ctx)
	if isWhiteListEnabled {
		whitelistedValidatorList, err := sd.pk.GetWhitelistedValidatorList(sdk.WrapSDKContext(ctx))
		if err != nil {
			return err
		}
		if whitelistedValidatorList != nil {
			//White List only applies to Standing Member Creation
			if createValidatorMsg.ValidatorType != stakingtypes.ValidatorTypeStanding {
				return nil
			} else {
				validatorAddress := createValidatorMsg.ValidatorAddress

				var isApprovedAddress = false
				for _, whitelistedVal := range whitelistedValidatorList {
					if whitelistedVal.ValidatorAddress == validatorAddress {

						if whitelistedVal.Moniker == createValidatorMsg.Description.Moniker {
							isApprovedAddress = true
							break
						} else {

							return sdkerrors.Wrap(
								permissionstypes.ErrNotMatchingMonikers,
								createValidatorMsg.Description.Moniker+" -- moniker does not match whitelisted validator's moniker",
							)
						}
					}
				}

				if isApprovedAddress {
					return nil
				} else {
					return sdkerrors.Wrap(
						permissionstypes.ErrUnauthorizedStandingMemberAddress,
						validatorAddress+" -- Unauthorized",
					)
				}
			}
		}
	}

	return nil

}

type StakingDelegationMessage struct {
	sk  vestingtypes.StakingKeeper
	cdc codec.BinaryCodec
	pk  permissionsmodulekeeper.Keeper
}

func NewStakingDelegationMessage(sk vestingtypes.StakingKeeper, cdc codec.BinaryCodec, pk permissionsmodulekeeper.Keeper) StakingDelegationMessage {
	return StakingDelegationMessage{
		sk:  sk,
		cdc: cdc,
		pk:  pk,
	}
}

func (sd StakingDelegationMessage) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *authz.MsgExec:
			// Check for bypassing authorization
			if err := sd.validateAuthz(ctx, msg); err != nil {
				return ctx, err
			}
		default:
			if err := sd.validateMsg(ctx, msg); err != nil {
				return ctx, err
			}
		}
	}

	return next(ctx, tx, simulate)
}

func (sd StakingDelegationMessage) validateAuthz(ctx sdk.Context, execMsg *authz.MsgExec) error {
	for _, v := range execMsg.Msgs {
		var innerMsg sdk.Msg
		if err := sd.cdc.UnpackAny(v, &innerMsg); err != nil {
			return sdkerrors.Wrap(err, "cannot unmarshal authz exec msgs")
		}
		if err := sd.validateMsg(ctx, innerMsg); err != nil {
			return err
		}
	}

	return nil
}

func (sd StakingDelegationMessage) validateMsg(ctx sdk.Context, msg sdk.Msg) error {

	delegationMessage, ok := msg.(*stakingtypes.MsgDelegate)
	if !ok {
		return nil
	}

	isWhiteListEnabled := sd.pk.GetIfExistsWhitelistEnabled(ctx)
	if isWhiteListEnabled {
		validatorAddress, err := sdk.ValAddressFromBech32(delegationMessage.ValidatorAddress)
		if err != nil {
			return err
		}
		delegatedValidator, _ := sd.sk.GetValidator(ctx, validatorAddress)

		if delegatedValidator.GetType() == stakingtypes.ValidatorTypeStanding {
			foundInWhitelist := sd.pk.FindValidator(ctx, validatorAddress)
			if !foundInWhitelist {
				fmt.Println("\n\n========================     StakingDelegationMessage     =============================")
				fmt.Println("----     NOT IN WHITELIST     ----")
				fmt.Println("validatorAddress - string: ", validatorAddress.String())
				fmt.Println("found in whitelist: ", foundInWhitelist)

				return permissionstypes.ErrInvalidDelegation
			}

		}
	}
	return nil

}
