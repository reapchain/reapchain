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
	KeyWhitelistEnabled                           = []byte("WhitelistEnabled")
	KeyPermissionsUnbondingTime                   = []byte("PermissionsUnbondingTime")
	KeyPermissionsMinimumInitialDepositEnabled    = []byte("PermissionsMinimumInitialDepositEnabled")
	KeyPermissionsMinimumInitialDepositPercentage = []byte("PermissionsMinimumInitialDepositPercentage")
)

var (
	DefaultWhitelistEnabled                                   = true
	DefaultPermissionsUnbondingTime                   string  = "4h"
	DefaultPermissionsMinimumInitialDepositEnabled    bool    = true
	DefaultPermissionsMinimumInitialDepositPercentage sdk.Dec = sdk.NewDec(1).Quo(sdk.NewDec(20))
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	whitelistEnabled bool,
	permissionsUnbondingTime string,
	permissionsMinimumInitialDepositEnabled bool,
	permissionsMinimumInitialDepositPercentage sdk.Dec,
) Params {
	return Params{
		WhitelistEnabled:                           whitelistEnabled,
		PermissionsUnbondingTime:                   permissionsUnbondingTime,
		PermissionsMinimumInitialDepositEnabled:    permissionsMinimumInitialDepositEnabled,
		PermissionsMinimumInitialDepositPercentage: permissionsMinimumInitialDepositPercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultWhitelistEnabled,
		DefaultPermissionsUnbondingTime,
		DefaultPermissionsMinimumInitialDepositEnabled,
		DefaultPermissionsMinimumInitialDepositPercentage,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyWhitelistEnabled, &p.WhitelistEnabled, validateWhitelistEnabled),
		paramtypes.NewParamSetPair(KeyPermissionsUnbondingTime, &p.PermissionsUnbondingTime, validatePermissionsUnbondingTime),
		paramtypes.NewParamSetPair(KeyPermissionsMinimumInitialDepositEnabled, &p.PermissionsMinimumInitialDepositEnabled, validateGovernanceMinimumInitialDepositEnabled),
		paramtypes.NewParamSetPair(KeyPermissionsMinimumInitialDepositPercentage, &p.PermissionsMinimumInitialDepositPercentage, validateGovernanceMinimumInitialDepositPercentage),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateWhitelistEnabled(p.WhitelistEnabled); err != nil {
		return err
	}

	if err := validatePermissionsUnbondingTime(p.PermissionsUnbondingTime); err != nil {
		return err
	}

	if err := validateGovernanceMinimumInitialDepositEnabled(p.PermissionsMinimumInitialDepositEnabled); err != nil {
		return err
	}

	if err := validateGovernanceMinimumInitialDepositPercentage(p.PermissionsMinimumInitialDepositPercentage); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
func validateWhitelistEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}

// validatePermissionsUnbondingTime validates the PermissionsUnbondingTime param
func validatePermissionsUnbondingTime(v interface{}) error {
	permissionsUnbondingTime, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	UnbondingTime, err := time.ParseDuration(permissionsUnbondingTime)
	if err != nil {
		return fmt.Errorf("invalid parameter type: %T", UnbondingTime)
	}

	return nil
}

// validateGovernanceMinimumInitialDepositEnabled validates the GovernanceMinimumInitialDepositEnabled param
func validateGovernanceMinimumInitialDepositEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return nil
}

// validateGovernanceMinimumInitialDepositPercentage validates the GovernanceMinimumInitialDepositPercentage param
func validateGovernanceMinimumInitialDepositPercentage(v interface{}) error {

	permissionsMinimumInitialDepositPercentage, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if permissionsMinimumInitialDepositPercentage.IsNegative() {
		return fmt.Errorf("min initial deposit cannot be negative: %s", v)
	}

	return nil
}
