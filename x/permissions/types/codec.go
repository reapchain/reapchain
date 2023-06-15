package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	cdctypes "github.com/reapchain/cosmos-sdk/codec/types"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	// this line is used by starport scaffolding # 1
	"github.com/reapchain/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterStandingMemberProposal{}, "permissions/RegisterStandingMember", nil)
	cdc.RegisterConcrete(&MsgRemoveStandingMemberProposal{}, "permissions/RemoveStandingMember", nil)
	cdc.RegisterConcrete(&MsgReplaceStandingMemberProposal{}, "permissions/ReplaceStandingMember", nil)

	// this line is used by starport scaffolding # 2
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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
