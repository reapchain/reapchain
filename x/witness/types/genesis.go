package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AnchoredList: []Anchored{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in anchored
	anchoredIdMap := make(map[uint64]bool)
	anchoredCount := gs.GetAnchoredCount()
	for _, elem := range gs.AnchoredList {
		if _, ok := anchoredIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for anchored")
		}
		if elem.Id >= anchoredCount {
			return fmt.Errorf("anchored id should be lower or equal than the last id")
		}
		anchoredIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
