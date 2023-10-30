package types

import (
	fmt "fmt"

	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
)

var (
	ParamStoreKeyEnableEscrow = []byte("EnableEscrow")
)

var _ paramtypes.ParamSet = &Params{}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(
	enableEscrow bool,
) Params {
	return Params{
		EnableEscrow: enableEscrow,
	}
}

func DefaultParams() Params {
	return Params{
		EnableEscrow: true,
	}
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyEnableEscrow, &p.EnableEscrow, validateBool),
	}
}

func (p Params) Validate() error { return nil }
