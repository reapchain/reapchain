package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EventConvertToNative = "convert_to_native"
	EventConvertToDenom  = "convert_to_denom"

	EventRegisterEscrowDenom           = "register_escrow_denom"
	EventRegisterEscrowDenomAndConvert = "register_escrow_denom_and_convert"

	EventToggleEscrowConversion    = "toggle_denom_conversion"
	EventAddToEscrowPool           = "add_to_escrow_pool"
	EventAddToEscrowPoolAndConvert = "add_to_escrow_pool_and_convert"

	AttributeKeyCosmosCoin = "cosmos_coin"
	AttributeKeyReceiver   = "receiver"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
