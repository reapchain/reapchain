package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	cdctypes "github.com/reapchain/cosmos-sdk/codec/types"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	"github.com/reapchain/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterStandingMemberProposal{}, "permissions/RegisterStandingMember", nil)
	cdc.RegisterConcrete(&MsgRemoveStandingMemberProposal{}, "permissions/RemoveStandingMember", nil)
	cdc.RegisterConcrete(&MsgReplaceStandingMemberProposal{}, "permissions/ReplaceStandingMember", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MsgRegisterStandingMemberProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MsgRemoveStandingMemberProposal{},
	)
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MsgReplaceStandingMemberProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
