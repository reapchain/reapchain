package types

import (
	"encoding/json"
	"fmt"
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
)

// NewGenesisState creates a new genesis state.
func NewGenesisState(params Params, registeredDenoms []RegisteredDenom, totalEscrowPool []sdk.Coin) GenesisState {
	return GenesisState{
		Params:           params,
		RegisteredDenoms: registeredDenoms,
		EscrowPools:      totalEscrowPool,
	}
}

// DefaultGenesisState sets default evm genesis state with empty accounts and
// default params and chain config values.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	seenDenom := make(map[string]bool)
	seenSupply := make(map[string]bool)

	for _, b := range gs.RegisteredDenoms {

		if seenDenom[b.Denom] {
			return fmt.Errorf("coin denomination duplicated on genesis for RegistereDenom: '%s'", b.Denom)
		}

		seenDenom[b.Denom] = true
	}

	for _, balance := range gs.EscrowPools {
		if seenSupply[balance.Denom] {
			return fmt.Errorf("coin denomination duplicated on genesis for EscrowPool: '%s'", balance.Denom)
		}
		seenSupply[balance.Denom] = true
	}

	return gs.Params.Validate()
}

func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
