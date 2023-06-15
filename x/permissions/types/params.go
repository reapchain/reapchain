package types

import (
	"fmt"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"time"

	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyForcedUnbondingTime            = []byte("ForcedUnbondingTime")
	DefaultForcedUnbondingTime string = "4h"
)

var (
	KeyGovernanceMinimumInitialDepositEnabled          = []byte("GovernanceMinimumInitialDepositEnabled")
	DefaultGovernanceMinimumInitialDepositEnabled bool = true
)

var (
	KeyGovernanceMinimumInitialDepositPercentage             = []byte("GovernanceMinimumInitialDepositPercentage")
	DefaultGovernanceMinimumInitialDepositPercentage sdk.Dec = sdk.NewDec(1).Quo(sdk.NewDec(20))
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	forcedUnbondingTime string,
	governanceMinimumInitialDepositEnabled bool,
	governanceMinimumInitialDepositPercentage sdk.Dec,
) Params {
	return Params{
		ForcedUnbondingTime:                       forcedUnbondingTime,
		GovernanceMinimumInitialDepositEnabled:    governanceMinimumInitialDepositEnabled,
		GovernanceMinimumInitialDepositPercentage: &governanceMinimumInitialDepositPercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultForcedUnbondingTime,
		DefaultGovernanceMinimumInitialDepositEnabled,
		DefaultGovernanceMinimumInitialDepositPercentage,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyForcedUnbondingTime, &p.ForcedUnbondingTime, validateForcedUnbondingTime),
		paramtypes.NewParamSetPair(KeyGovernanceMinimumInitialDepositEnabled, &p.GovernanceMinimumInitialDepositEnabled, validateGovernanceMinimumInitialDepositEnabled),
		paramtypes.NewParamSetPair(KeyGovernanceMinimumInitialDepositPercentage, &p.GovernanceMinimumInitialDepositPercentage, validateGovernanceMinimumInitialDepositPercentage),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateForcedUnbondingTime(p.ForcedUnbondingTime); err != nil {
		return err
	}

	if err := validateGovernanceMinimumInitialDepositEnabled(p.GovernanceMinimumInitialDepositEnabled); err != nil {
		return err
	}

	if err := validateGovernanceMinimumInitialDepositPercentage(p.GovernanceMinimumInitialDepositPercentage); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateForcedUnbondingTime validates the ForcedUnbondingTime param
func validateForcedUnbondingTime(v interface{}) error {
	forcedUnbondingTime, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	UnbondingTime, err := time.ParseDuration(forcedUnbondingTime)
	if err != nil {
		return fmt.Errorf("invalid parameter type: %T", UnbondingTime)
	}

	return nil
}

// validateGovernanceMinimumInitialDepositEnabled validates the GovernanceMinimumInitialDepositEnabled param
func validateGovernanceMinimumInitialDepositEnabled(v interface{}) error {
	governanceMinimumInitialDepositEnabled, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = governanceMinimumInitialDepositEnabled

	return nil
}

// validateGovernanceMinimumInitialDepositPercentage validates the GovernanceMinimumInitialDepositPercentage param
func validateGovernanceMinimumInitialDepositPercentage(v interface{}) error {
	governanceMinimumInitialDepositPercentage, ok := v.(*sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if governanceMinimumInitialDepositPercentage.IsNegative() {
		return fmt.Errorf("min initial deposit cannot be negative: %s", v)
	}

	return nil
}
