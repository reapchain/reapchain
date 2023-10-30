package keeper

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

func (k Keeper) RegisterEscrowDenom(
	ctx sdk.Context,
	denom string,
	initialPoolAmount sdk.Int,
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
	k.HandleAddToEscrowPool(ctx, denom, initialPoolAmount)

	return &newRegisteredDenom, nil
}

func (k Keeper) RegisterEscrowDenomAndConvert(
	ctx sdk.Context,
	denom string,
	initialPoolAmount sdk.Int,
	from string,
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
	k.HandleAddToEscrowPool(ctx, denom, initialPoolAmount)

	fromAddr, err := sdk.AccAddressFromBech32(from)
	currentNativeBalance := k.bankKeeper.GetBalance(ctx, fromAddr, sdk.DefaultBondDenom)

	escrowSupply, found := k.GetEscrowPoolByDenom(ctx, denom)
	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrDenomNotFound, "denom not registered: %s", denom,
		)
	}

	if initialPoolAmount.GT(currentNativeBalance.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "not enough balance in account for denom: %s", denom,
		)
	}

	burnCoins := sdk.Coins{sdk.Coin{
		Denom:  sdk.DefaultBondDenom,
		Amount: initialPoolAmount,
	}}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, fromAddr, types.ModuleName, burnCoins)
	if err != nil {
		return nil, err
	}
	burnCoinsErr := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if burnCoinsErr != nil {
		return nil, burnCoinsErr
	}

	mintCoins := sdk.Coins{sdk.Coin{
		Denom:  denom,
		Amount: initialPoolAmount,
	}}

	mintCoinsErr := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if mintCoinsErr != nil {
		return nil, mintCoinsErr
	}
	sendCoinsFromModuleToAccErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fromAddr, mintCoins)
	if sendCoinsFromModuleToAccErr != nil {
		return nil, sendCoinsFromModuleToAccErr
	}

	newEscrowSupplyCoins := sdk.Coins{sdk.Coin{
		Denom:  denom,
		Amount: escrowSupply.Coins.AmountOf(denom).Sub(initialPoolAmount),
	}}

	newEscrowSupply := types.EscrowPool{
		Denom: denom,
		Coins: newEscrowSupplyCoins,
	}

	k.SetEscrowPool(ctx, newEscrowSupply)

	return &newRegisteredDenom, nil
}

func (k Keeper) HandleAddToEscrowPool(
	ctx sdk.Context,
	denom string,
	amount sdk.Int,
) (*types.EscrowPool, error) {
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

	escrowSupply, found := k.GetEscrowPoolByDenom(ctx, denom)

	if !found {
		newSupply := sdk.Coins{sdk.Coin{Denom: denom, Amount: amount}}

		newEscrowSupply := types.EscrowPool{
			Denom: denom,
			Coins: newSupply,
		}
		k.SetEscrowPool(ctx, newEscrowSupply)
		return &newEscrowSupply, nil
	}

	if found {

		additionalSupply := sdk.Coin{Denom: denom, Amount: amount}
		currentSupply := escrowSupply.Coins

		newSupply := currentSupply.Add(additionalSupply)
		newEscrowSupply := types.EscrowPool{
			Denom: denom,
			Coins: newSupply,
		}

		k.SetEscrowPool(ctx, newEscrowSupply)
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
