package escrow

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	"github.com/reapchain/reapchain/v8/x/escrow/keeper"
	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// NewErc20ProposalHandler creates a governance handler to manage new proposal types.
func NewEscrowProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterEscrowDenomProposal:
			return handleRegisterEscrowDenomProposal(ctx, k, c)
		case *types.RegisterEscrowDenomAndConvertProposal:
			return handleRegisterEscrowDenomAndConvertProposal(ctx, k, c)
		case *types.ToggleEscrowConversionProposal:
			return handleToggleEscrowConversionProposal(ctx, k, c)
		case *types.AddToEscrowPoolProposal:
			return handleAddToEscrowPoolProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterEscrowDenomAndConvertProposal(ctx sdk.Context, k *keeper.Keeper, redp *types.RegisterEscrowDenomAndConvertProposal) error {
	registeredDenom, err := k.RegisterEscrowDenomAndConvert(ctx, redp.Denom, redp.InitialPoolBalance, redp.Proposer)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterEscrowDenomAndConvert,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, registeredDenom.Denom),
		),
	)

	return nil
}

func handleRegisterEscrowDenomProposal(ctx sdk.Context, k *keeper.Keeper, redp *types.RegisterEscrowDenomProposal) error {
	registeredDenom, err := k.RegisterEscrowDenom(ctx, redp.Denom, redp.InitialPoolBalance)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterEscrowDenom,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, registeredDenom.Denom),
		),
	)

	return nil
}

func handleToggleEscrowConversionProposal(ctx sdk.Context, k *keeper.Keeper, tdcp *types.ToggleEscrowConversionProposal) error {
	registeredDenom, err := k.ToggleEscrowConversion(ctx, tdcp.Denom)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventToggleEscrowConversion,
			sdk.NewAttribute(types.AttributeKeyCosmosCoin, registeredDenom.Denom),
		),
	)

	return nil
}

func handleAddToEscrowPoolProposal(ctx sdk.Context, k *keeper.Keeper, aesp *types.AddToEscrowPoolProposal) error {
	escrowPool, err := k.HandleAddToEscrowPool(ctx, aesp.Denom, aesp.Amount)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddToEscrowPool,
			sdk.NewAttribute(types.AttributeKeyReceiver, escrowPool.Denom),
		),
	)

	return nil
}
