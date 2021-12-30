package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MintCoins{}, "reapchain/MintCoins", nil)
	cdc.RegisterConcrete(&BurnCoins{}, "reapchain/BurnCoins", nil)
	cdc.RegisterConcrete(&TransferMintedCoins{}, "reapchain/TransferMintedCoins", nil)
	cdc.RegisterConcrete(&BurnCoinsFromAccount{}, "reapchain/BurnCoinsFromAccount", nil)
} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MintCoins{}, &BurnCoins{}, &TransferMintedCoins{}, &BurnCoinsFromAccount{},
	)
}

var (
	Amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
