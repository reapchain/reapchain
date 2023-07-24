package types

import (
	"fmt"
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyPodcWhitelistEnabled           = []byte("PodcWhitelistEnabled")
	KeyGovMinInitialDepositEnabled    = []byte("GovMinInitialDepositEnabled")
	KeyGovMinInitialDepositPercentage = []byte("GovMinInitialDepositPercentage")
)

var (
	DefaultWhitelistEnabled                       = true
	DefaultGovMinInitialDepositEnabled    bool    = true
	DefaultGovMinInitialDepositPercentage sdk.Dec = sdk.NewDec(1).Quo(sdk.NewDec(20))
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	PodcWhitelistEnabled bool,
	GovMinInitialDepositEnabled bool,
	GovMinInitialDepositPercentage sdk.Dec,
) Params {
	return Params{
		PodcWhitelistEnabled:           PodcWhitelistEnabled,
		GovMinInitialDepositEnabled:    GovMinInitialDepositEnabled,
		GovMinInitialDepositPercentage: GovMinInitialDepositPercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultWhitelistEnabled,
		DefaultGovMinInitialDepositEnabled,
		DefaultGovMinInitialDepositPercentage,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPodcWhitelistEnabled, &p.PodcWhitelistEnabled, validateWhitelistEnabled),
		paramtypes.NewParamSetPair(KeyGovMinInitialDepositEnabled, &p.GovMinInitialDepositEnabled, validateGovMinInitialDepositEnabled),
		paramtypes.NewParamSetPair(KeyGovMinInitialDepositPercentage, &p.GovMinInitialDepositPercentage, validateGovMinInitialDepositPercentage),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateWhitelistEnabled(p.PodcWhitelistEnabled); err != nil {
		return err
	}

	if err := validateGovMinInitialDepositEnabled(p.GovMinInitialDepositEnabled); err != nil {
		return err
	}

	if err := validateGovMinInitialDepositPercentage(p.GovMinInitialDepositPercentage); err != nil {
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

// validateGovernanceMinimumInitialDepositEnabled validates the GovernanceMinimumInitialDepositEnabled param
func validateGovMinInitialDepositEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	return nil
}

// validateGovernanceMinimumInitialDepositPercentage validates the GovernanceMinimumInitialDepositPercentage param
func validateGovMinInitialDepositPercentage(v interface{}) error {

	govMinInitialDepositPercentage, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}
	if govMinInitialDepositPercentage.IsNegative() {
		return fmt.Errorf("min initial deposit cannot be negative: %s", v)
	}

	return nil
}
