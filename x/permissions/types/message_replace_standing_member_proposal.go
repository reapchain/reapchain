package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

const TypeMsgReplaceStandingMemberProposal = "replace_standing_member_proposal"

var _ govtypes.Content = &MsgReplaceStandingMemberProposal{}

func init() {
	govtypes.RegisterProposalType(TypeMsgReplaceStandingMemberProposal)
	govtypes.RegisterProposalTypeCodec(&MsgReplaceStandingMemberProposal{}, "permissions/MsgReplaceStandingMemberProposal")

}
func NewMsgReplaceStandingMemberProposal(
	title string,
	description string,
	existingValAddr string,
	replacementValAddr string,
	replacementAccAddr string,
	moniker string) *MsgReplaceStandingMemberProposal {
	return &MsgReplaceStandingMemberProposal{
		Title:                       title,
		Description:                 description,
		ExistingValidatorAddress:    existingValAddr,
		ReplacementValidatorAddress: replacementValAddr,
		ReplacementAccountAddress:   replacementAccAddr,
		ReplacementMoniker:          moniker,
	}
}
func (msg MsgReplaceStandingMemberProposal) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.ReplacementAccountAddress)
	return []sdk.AccAddress{addr}
}

// ProposalRoute returns router key for this proposal
func (*MsgReplaceStandingMemberProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*MsgReplaceStandingMemberProposal) ProposalType() string {
	return TypeMsgReplaceStandingMemberProposal
}

func (msg *MsgReplaceStandingMemberProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReplaceStandingMemberProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ReplacementAccountAddress)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
