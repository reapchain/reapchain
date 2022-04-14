package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &TransferMintedCoins{}


func NewMsgTransferMintedCoins(creator string, amount string, account string) *TransferMintedCoins {
	return &TransferMintedCoins{
		Creator: creator,
		Amount: amount,
		Account: account,
	}
}
// Route ...
func (msg TransferMintedCoins) Route() string {
	return RouterKey
}

// Type ...
func (msg TransferMintedCoins) Type() string {
	return "TransferMintedCoins"
}

// GetSigners ...
func (msg *TransferMintedCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes ...
func (msg *TransferMintedCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic ...
func (msg *TransferMintedCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	//_, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}