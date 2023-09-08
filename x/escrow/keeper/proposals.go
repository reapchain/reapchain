package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func (k Keeper) RegisterEscrowDenom(
	ctx sdk.Context,
	denom string,
	initialSupply sdk.Int,
) (*types.RegisteredDenom, error) {
	params := k.GetParams(ctx)
	if !params.EnableEscrow {
		return nil, sdkerrors.Wrap(
			types.ErrEscrowDisabled, "registration is currently disabled by governance",
		)
	}

	if k.IsDenomRegistered(ctx, denom) {
		return nil, sdkerrors.Wrapf(
			types.ErrDenomAlreadyExists, "coin denomination already registered: %s", denom,
		)
	}

	newRegisteredDenom := NewEscrowDenom(denom, true)
	k.RegisterDenom(ctx, newRegisteredDenom)
	k.HandleAddEscrowSupply(ctx, denom, initialSupply)
	return &newRegisteredDenom, nil
}

func (k Keeper) HandleAddEscrowSupply(
	ctx sdk.Context,
	denom string,
	amount sdk.Int,
) (*types.EscrowSupply, error) {
	// Check if the conversion is globally enabled
	params := k.GetParams(ctx)
	if !params.EnableEscrow {
		return nil, sdkerrors.Wrap(
			types.ErrEscrowDisabled, "registration is currently disabled by governance",
		)
	}

	if !k.IsDenomRegistered(ctx, denom) {
		return nil, sdkerrors.Wrapf(
			types.ErrDenomNotFound, "denom not registered: %s", denom,
		)
	}

	escrowSupply, found := k.GetEscrowSupplyByDenom(ctx, denom)

	if !found {
		newSupply := sdk.Coins{sdk.Coin{Denom: denom, Amount: amount}}

		newEscrowSupply := types.EscrowSupply{
			Denom: denom,
			Coins: newSupply,
		}
		k.SetEscrowSupply(ctx, newEscrowSupply)
		return &newEscrowSupply, nil
	}

	if found {

		additionalSupply := sdk.Coin{Denom: denom, Amount: amount}
		currentSupply := escrowSupply.Coins

		newSupply := currentSupply.Add(additionalSupply)
		newEscrowSupply := types.EscrowSupply{
			Denom: denom,
			Coins: newSupply,
		}

		k.SetEscrowSupply(ctx, newEscrowSupply)
		return &newEscrowSupply, nil
	}

	return &escrowSupply, nil
}

func (k Keeper) ToggleEscrowConversion(
	ctx sdk.Context,
	denom string,
) (types.RegisteredDenom, error) {

	params := k.GetParams(ctx)
	if !params.EnableEscrow {
		return types.RegisteredDenom{}, sdkerrors.Wrap(
			types.ErrEscrowDisabled, "registration is currently disabled by governance",
		)
	}
	registeredDenom, found := k.GetRegisteredDenom(ctx, denom)
	if !found {
		return types.RegisteredDenom{}, sdkerrors.Wrapf(
			types.ErrDenomNotFound, "denom '%s' not registered by id", denom,
		)
	}

	registeredDenom.Enabled = !registeredDenom.Enabled

	k.SetRegisteredDenom(ctx, registeredDenom)
	return registeredDenom, nil
}
