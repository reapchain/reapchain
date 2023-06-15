package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

const TypeMsgRegisterStandingMemberProposal = "MsgRegisterStandingMember"

var _ govtypes.Content = &MsgRegisterStandingMemberProposal{}

func init() {
	govtypes.RegisterProposalType(TypeMsgRegisterStandingMemberProposal)
	govtypes.RegisterProposalTypeCodec(&MsgRegisterStandingMemberProposal{}, "permissions/MsgRegisterStandingMemberProposal")

}

func NewMsgRegisterStandingMemberProposal(title string, description string, validatorAddress string, accountAddress string, moniker string) *MsgRegisterStandingMemberProposal {
	return &MsgRegisterStandingMemberProposal{

		Title:            title,
		Description:      description,
		ValidatorAddress: validatorAddress,
		AccountAddress:   accountAddress,
		Moniker:          moniker,
	}
}

// ProposalRoute returns router key for this proposal
func (*MsgRegisterStandingMemberProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (*MsgRegisterStandingMemberProposal) ProposalType() string {
	return TypeMsgRegisterStandingMemberProposal
}

func (msg *MsgRegisterStandingMemberProposal) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.AccountAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgRegisterStandingMemberProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterStandingMemberProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AccountAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
