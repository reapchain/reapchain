package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

const TypeMsgRemoveStandingMemberProposal = "remove_standing_member_proposal"

var _ govtypes.Content = &MsgRemoveStandingMemberProposal{}

func init() {
	govtypes.RegisterProposalType(TypeMsgRemoveStandingMemberProposal)
	govtypes.RegisterProposalTypeCodec(&MsgRemoveStandingMemberProposal{}, "permissions/MsgRemoveStandingMemberProposal")

}

func NewMsgRemoveStandingMemberProposal(title string, description string, validatorAddress string) *MsgRemoveStandingMemberProposal {
	return &MsgRemoveStandingMemberProposal{
		Title:            title,
		Description:      description,
		ValidatorAddress: validatorAddress,
	}
}

// ProposalRoute returns router key for this proposal
func (*MsgRemoveStandingMemberProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*MsgRemoveStandingMemberProposal) ProposalType() string {
	return TypeMsgRemoveStandingMemberProposal
}

func (msg *MsgRemoveStandingMemberProposal) GetSigners() []sdk.AccAddress {
	sender, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(sender)}
}

func (msg *MsgRemoveStandingMemberProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveStandingMemberProposal) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
