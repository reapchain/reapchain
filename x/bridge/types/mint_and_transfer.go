package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

const TypeMsgMintAndTransfer = "mint_and_transfer"

var _ sdk.Msg = &MsgMintAndTransfer{}

func NewMsgMintAndTransfer(authorizedAccount string, address string, amount string) *MsgMintAndTransfer {
	return &MsgMintAndTransfer{
		AuthorizedAccount: authorizedAccount,
		Address:           address,
		Amount:            amount,
	}
}

func (msg *MsgMintAndTransfer) Route() string {
	return RouterKey
}

func (msg *MsgMintAndTransfer) Type() string {
	return TypeMsgMintAndTransfer
}

func (msg *MsgMintAndTransfer) GetSigners() []sdk.AccAddress {
	authorizedAccount, err := sdk.AccAddressFromBech32(msg.AuthorizedAccount)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authorizedAccount}
}

func (msg *MsgMintAndTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintAndTransfer) ValidateBasic() error {

	coins, coinParseError := sdk.ParseCoinNormalized(msg.Amount)
	if coinParseError != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "send amount is invalid: %s", coinParseError)
	}

	if coins.GetDenom() != sdk.DefaultBondDenom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "send denomination is invalid")
	}

	_, err := sdk.AccAddressFromBech32(msg.AuthorizedAccount)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authorizedAccount address (%s)", err)
	}

	if msg.AuthorizedAccount != AuthorizedAddress {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "address does not have permission")
	}

	_, addressParseError := sdk.AccAddressFromBech32(msg.Address)
	if addressParseError != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", addressParseError)
	}

	return nil
}
