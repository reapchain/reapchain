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

	"time"
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
	whitelistedValidatorList, err := sd.pk.GetWhitelistedValidatorList(sdk.WrapSDKContext(ctx))
	if err != nil {
		return err
	}

	if whitelistedValidatorList != nil {
		//White List only applies to Standing Member Creation
		if createValidatorMsg.ValidatorType != "standing" {
			return nil
		} else {
			validatorAddress := createValidatorMsg.ValidatorAddress

			var isApprovedAddress = false
			for _, whitelistedVal := range whitelistedValidatorList {
				if whitelistedVal.ValidatorAddress == validatorAddress {
					isApprovedAddress = true
					break
				}
			}

			fmt.Println("\n=================================================")
			fmt.Println("STAKING ANTE HANDLER - validateMsg", time.Now().Format(time.RFC822))
			fmt.Println("ADDRESS: ", validatorAddress)
			fmt.Println("isApprovedAddress: ", isApprovedAddress)
			fmt.Println("=================================================")

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
	return nil

}
