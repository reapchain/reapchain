package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MintCoins{}


func NewMsgMintCoins(creator string, amount string) *MintCoins {
	return &MintCoins{
		Creator: creator,
		Amount: amount,
	}
}
// Route ...
func (msg MintCoins) Route() string {
	return RouterKey
}

// Type ...
func (msg MintCoins) Type() string {
	return "MintCoins"
}

// GetSigners ...
func (msg *MintCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *MintCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *MintCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	//_, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}