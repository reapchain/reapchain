package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	ethermint "github.com/reapchain/ethermint/types"
)

var (
	_ sdk.Msg = &MsgConvertToNative{}
	_ sdk.Msg = &MsgConvertToDenom{}
)

const (
	TypeMsgConvertToNative = "convert_to_native"
	TypeMsgConvertToDenom  = "convert_to_denom"
)

// NewMsgConvertToNative creates a new instance of MsgConvertToNative
func NewMsgConvertToNative(coin sdk.Coin, sender sdk.AccAddress) *MsgConvertToNative { // nolint: interfacer
	return &MsgConvertToNative{
		Coin:   coin,
		Sender: sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertToNative) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertToNative) Type() string { return TypeMsgConvertToNative }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertToNative) ValidateBasic() error {

	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertToNative) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertToNative) GetSigners() []sdk.AccAddress {
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// NewMsgConvertToDenom creates a new instance of MsgConvertToDenom
func NewMsgConvertToDenom(amount sdk.Int, denom string, sender sdk.AccAddress) *MsgConvertToDenom { // nolint: interfacer
	return &MsgConvertToDenom{
		Amount: amount,
		Denom:  denom,
		Sender: sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgConvertToDenom) Route() string { return RouterKey }

// Type should return the action
func (msg MsgConvertToDenom) Type() string { return TypeMsgConvertToDenom }

// ValidateBasic runs stateless checks on the message
func (msg MsgConvertToDenom) ValidateBasic() error {

	EVMAddress := removePrefix(msg.Denom)
	if err := ethermint.ValidateAddress(EVMAddress); err != nil {
		return sdkerrors.Wrap(err, "Invalid EVM address")
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgConvertToDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgConvertToDenom) GetSigners() []sdk.AccAddress {
	addr := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}
