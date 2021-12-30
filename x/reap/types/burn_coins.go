package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &BurnCoins{}


func NewMsgBurnCoins(creator string, amount string) *BurnCoins {
	return &BurnCoins{
		Creator: creator,
		Amount: amount,
	}
}
// Route ...
func (msg BurnCoins) Route() string {
	return RouterKey
}

// Type ...
func (msg BurnCoins) Type() string {
	return "BurnCoins"
}

// GetSigners ...
func (msg *BurnCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *BurnCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *BurnCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	//_, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}



var _ sdk.Msg = &BurnCoinsFromAccount{}


func NewMsgBurnCoinsFromAccount(creator string, amount string, account string) *BurnCoinsFromAccount {
	return &BurnCoinsFromAccount{
		Creator: creator,
		Amount: amount,
		Account: account,
	}
}
// Route ...
func (msg BurnCoinsFromAccount) Route() string {
	return RouterKey
}

// Type ...
func (msg BurnCoinsFromAccount) Type() string {
	return "BurnCoinsFromAccount"
}

// GetSigners ...
func (msg *BurnCoinsFromAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *BurnCoinsFromAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *BurnCoinsFromAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	//_, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}