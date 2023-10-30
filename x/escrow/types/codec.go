package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/msgservice"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
)

var (
	amino = codec.NewLegacyAmino()

	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	AminoCdc = codec.NewAminoCodec(amino)
)

const (
	convertToNative = "reapchain/MsgConvertToNative"
	convertToDenom  = "reapchain/MsgConvertToDenom"
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConvertToNative{},
		&MsgConvertToDenom{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&RegisterEscrowDenomProposal{},
		&RegisterEscrowDenomAndConvertProposal{},
		&ToggleEscrowConversionProposal{},
		&AddToEscrowPoolProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgConvertToNative{}, convertToNative, nil)
	cdc.RegisterConcrete(&MsgConvertToDenom{}, convertToDenom, nil)
}
