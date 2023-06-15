package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/address"
)

var (
	ValidatorsKey = []byte{0x21} // prefix for each key to a validator

)

// GetValidatorKey creates the key for the validator with address
// VALUE: staking/Validator
func GetValidatorKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, address.MustLengthPrefix(operatorAddr)...)
}
