package keeper

import (
	"context"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

var _ types.MsgServer = &Keeper{}

func (k Keeper) ConvertToNative(
	goCtx context.Context,
	msg *types.MsgConvertToNative,
) (*types.MsgConvertToNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)

	escrowParams := k.GetParams(ctx)
	if !escrowParams.EnableEscrow {
		return nil, sdkerrors.Wrap(types.ErrEscrowDisabled, "module is currently disabled by governance")

	}

	registeredDenom, found := k.GetRegisteredDenom(ctx, msg.Coin.Denom)

	if !registeredDenom.Enabled {
		return nil, sdkerrors.Wrap(types.ErrDenominationDisabled, registeredDenom.Denom)
	}

	if !found {
		return nil, sdkerrors.Wrap(types.ErrDenomNotFound, registeredDenom.Denom)
	}

	currentDenomBalance := k.bankKeeper.GetBalance(ctx, sender, msg.Coin.Denom)

	if msg.Coin.Amount.GT(currentDenomBalance.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "not enough balance in account for denom: %s", msg.Coin.Denom,
		)
	}
	burnCoins := sdk.Coins{msg.Coin}
	sendFromAccToModuleErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnCoins)
	if sendFromAccToModuleErr != nil {
		return nil, sendFromAccToModuleErr
	}

	burnCoinsErr := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if burnCoinsErr != nil {
		return nil, burnCoinsErr
	}

	mintCoins := sdk.Coins{sdk.Coin{
		Denom:  sdk.DefaultBondDenom,
		Amount: msg.Coin.Amount,
	}}
	mintCoinsErr := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if mintCoinsErr != nil {
		return nil, mintCoinsErr
	}
	sendCoinsFromModuleToAccErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, mintCoins)
	if sendCoinsFromModuleToAccErr != nil {
		return nil, sendCoinsFromModuleToAccErr
	}
	escrowPool, _ := k.GetEscrowPoolByDenom(ctx, msg.Coin.Denom)

	newEscrowPoolAmount := sdk.Coin{
		Denom:  msg.Coin.Denom,
		Amount: escrowPool.Amount.Add(msg.Coin.Amount),
	}
	k.SetEscrowPool(ctx, newEscrowPoolAmount)

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventConvertToNative,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Coin.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Coin.Denom),
			),
		},
	)

	return &types.MsgConvertToNativeResponse{}, nil

}

func (k Keeper) ConvertToDenom(
	goCtx context.Context,
	msg *types.MsgConvertToDenom,
) (*types.MsgConvertToDenomResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)

	escrowParams := k.GetParams(ctx)
	if !escrowParams.EnableEscrow {
		return nil, sdkerrors.Wrap(types.ErrEscrowDisabled, "module is currently disabled by governance")

	}

	registeredDenom, found := k.GetRegisteredDenom(ctx, msg.Denom)

	if !registeredDenom.Enabled {
		return nil, sdkerrors.Wrap(types.ErrDenominationDisabled, registeredDenom.Denom)
	}

	if !found {
		return nil, sdkerrors.Wrap(types.ErrDenomNotFound, registeredDenom.Denom)
	}

	currentNativeBalance := k.bankKeeper.GetBalance(ctx, sender, sdk.DefaultBondDenom)

	if msg.Amount.GT(currentNativeBalance.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "not enough balance in account for denom: %s", msg.Denom,
		)
	}

	escrowPool, found := k.GetEscrowPoolByDenom(ctx, msg.Denom)

	if !found {
		return nil, sdkerrors.Wrapf(
			types.ErrDenomNotFound, "denom not registered: %s", msg.Denom,
		)
	}

	if escrowPool.Amount.LT(msg.Amount) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds, "not enough escrow supply for denom: %s", msg.Denom,
		)
	}

	burnCoins := sdk.Coins{sdk.Coin{
		Denom:  sdk.DefaultBondDenom,
		Amount: msg.Amount,
	}}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnCoins)
	if err != nil {
		return nil, err
	}
	burnCoinsErr := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if burnCoinsErr != nil {
		return nil, burnCoinsErr
	}

	mintCoins := sdk.Coins{sdk.Coin{
		Denom:  msg.Denom,
		Amount: msg.Amount,
	}}

	mintCoinsErr := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if mintCoinsErr != nil {
		return nil, mintCoinsErr
	}
	sendCoinsFromModuleToAccErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, mintCoins)
	if sendCoinsFromModuleToAccErr != nil {
		return nil, sendCoinsFromModuleToAccErr
	}

	newEscrowPoolAmount := sdk.Coin{
		Denom:  msg.Denom,
		Amount: escrowPool.Amount.Sub(msg.Amount),
	}

	k.SetEscrowPool(ctx, newEscrowPoolAmount)
	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventConvertToDenom,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
				sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.String()),
				sdk.NewAttribute(types.AttributeKeyCosmosCoin, msg.Denom),
			),
		},
	)

	return &types.MsgConvertToDenomResponse{}, nil

}
