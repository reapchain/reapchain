package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/address"
	"time"
)

var (
	ValidatorsKey               = []byte{0x21} // prefix for each key to a validator
	DefaultUnbondingTime        = "4s"
	DefaultRemovalBlockInterval = int64(1)
	DefaultSafeTimeInterval     = 15 * time.Second
)

// GetValidatorKey creates the key for the validator with address
// VALUE: staking/Validator
func GetValidatorKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, address.MustLengthPrefix(operatorAddr)...)
}
