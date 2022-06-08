package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/msgservice"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"
	"github.com/reapchain/cosmos-sdk/x/auth/vesting/exported"
	sdkvesting "github.com/reapchain/cosmos-sdk/x/auth/vesting/types"
)

// ModuleCdc references the global erc20 module codec. Note, the codec should
// ONLY be used in certain instances of tests and for JSON encoding.
//
// The actual codec used for serialization should be provided to modules/erc20 and
// defined at the application level.
var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// RegisterInterface associates protoName with AccountI and VestingAccount
// Interfaces and creates a registry of it's concrete implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"cosmos.vesting.v1beta1.VestingAccount",
		(*exported.VestingAccount)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*authtypes.GenesisAccount)(nil),
		&sdkvesting.ContinuousVestingAccount{},
		&sdkvesting.DelayedVestingAccount{},
		&sdkvesting.PeriodicVestingAccount{},
		&sdkvesting.PermanentLockedAccount{},
		&ClawbackVestingAccount{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgClawback{},
		&MsgCreateClawbackVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
