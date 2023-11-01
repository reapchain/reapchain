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

	escrowPool, found := k.GetEscrowPoolByDenom(ctx, denom)
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

	newEscrowPoolBalance := sdk.Coin{
		Denom:  denom,
		Amount: escrowPool.Amount.Sub(initialPoolAmount),
	}

	k.SetEscrowPool(ctx, newEscrowPoolBalance)

	return &newRegisteredDenom, nil
}

func (k Keeper) HandleAddToEscrowPool(
	ctx sdk.Context,
	denom string,
	amount sdk.Int,
) (*sdk.Coin, error) {
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

	escrowPool, found := k.GetEscrowPoolByDenom(ctx, denom)

	if !found {
		newEscrowPoolAmount := sdk.Coin{Denom: denom, Amount: amount}
		//newEscrowPoolAmount := types.EscrowPool{Balance: newPoolAmount}

		k.SetEscrowPool(ctx, newEscrowPoolAmount)
		return &newEscrowPoolAmount, nil
	}

	if found {

		additionalPoolAmount := sdk.Coin{Denom: denom, Amount: amount}
		currentPoolAmount := escrowPool

		newEscrowPoolAmount := currentPoolAmount.Add(additionalPoolAmount)
		k.SetEscrowPool(ctx, newEscrowPoolAmount)
		return &newEscrowPoolAmount, nil
	}

	return &escrowPool, nil
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
