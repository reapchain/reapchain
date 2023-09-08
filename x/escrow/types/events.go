package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EventConvertToNative = "convert_to_native"
	EventConvertToDenom  = "convert_to_denom"

	EventRegisterEscrowDenom    = "register_escrow_denom"
	EventToggleEscrowConversion = "toggle_denom_conversion"
	EventAddEscrowSupply        = "add_escrow_supply"

	AttributeKeyCosmosCoin = "cosmos_coin"
	AttributeKeyReceiver   = "receiver"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
