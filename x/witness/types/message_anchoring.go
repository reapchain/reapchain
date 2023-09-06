package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
)

const TypeMsgAnchoring = "anchoring"

var _ sdk.Msg = &MsgAnchoring{}

func NewMsgAnchoring(creator string, blockHash string, blockHeight string, chainID string, filter string) *MsgAnchoring {
	return &MsgAnchoring{
		Creator:     creator,
		BlockHash:   blockHash,
		BlockHeight: blockHeight,
		ChainID:     chainID,
		Filter:      filter,
	}
}

func (msg *MsgAnchoring) Route() string {
	return RouterKey
}

func (msg *MsgAnchoring) Type() string {
	return TypeMsgAnchoring
}

func (msg *MsgAnchoring) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAnchoring) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAnchoring) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
