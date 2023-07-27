package ante

import (
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/x/authz"
	govkeeper "github.com/reapchain/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
	permissionsmodulekeeper "github.com/reapchain/reapchain/v8/x/permissions/keeper"
	permissionsmoduletypes "github.com/reapchain/reapchain/v8/x/permissions/types"
)

type GovernanceSubmitProposalMessage struct {
	cdc codec.BinaryCodec
	pk  permissionsmodulekeeper.Keeper
	gk  govkeeper.Keeper
}

func NewGovernanceSubmitProposalMessage(cdc codec.BinaryCodec, pk permissionsmodulekeeper.Keeper, gk govkeeper.Keeper) GovernanceSubmitProposalMessage {
	return GovernanceSubmitProposalMessage{
		cdc: cdc,
		pk:  pk,
		gk:  gk,
	}
}

func (sd GovernanceSubmitProposalMessage) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
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

func (sd GovernanceSubmitProposalMessage) validateAuthz(ctx sdk.Context, execMsg *authz.MsgExec) error {
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

func (sd GovernanceSubmitProposalMessage) validateMsg(ctx sdk.Context, msg sdk.Msg) error {

	govSubmitProposalMsg, ok := msg.(*govtypes.MsgSubmitProposal)
	if !ok {
		return nil
	}

	govModuleParams := sd.gk.GetDepositParams(ctx)
	permissionsModuleParams := sd.pk.GetParams(ctx)

	minimumInitialDepositEnabled := permissionsModuleParams.GovMinInitialDepositEnabled

	if minimumInitialDepositEnabled {

		govSubmitProposalMsg.GetInitialDeposit()
		minimumInitialDepositPercentage := permissionsModuleParams.GovMinInitialDepositPercentage
		govParamMinimumDeposit := govModuleParams.MinDeposit.AmountOf(sdk.DefaultBondDenom)
		requiredInitialDeposit := minimumInitialDepositPercentage.MulInt(govParamMinimumDeposit).RoundInt().BigInt()
		initialDeposit := govSubmitProposalMsg.InitialDeposit.AmountOf(sdk.DefaultBondDenom).BigInt()
		if initialDeposit.Cmp(requiredInitialDeposit) == -1 {

			return sdkerrors.Wrap(
				permissionsmoduletypes.ErrInsufficientInitialDeposit,
				initialDeposit.String()+" -- deposit is less than "+requiredInitialDeposit.String()+sdk.DefaultBondDenom,
			)
		}
	}

	return nil

}
