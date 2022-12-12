package types

import (
	"testing"
	"time"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestValidateGenesis() {
	newGen := NewGenesisState(DefaultParams(), []Incentive{}, []GasMeter{})

	testCases := []struct {
		name     string
		genState *GenesisState
		expPass  bool
	}{
		{
			"valid genesis constructor",
			&newGen,
			true,
		},
		{
			"default",
			DefaultGenesisState(),
			true,
		},
		{
			"valid genesis",
			&GenesisState{
				Params:     DefaultParams(),
				Incentives: []Incentive{},
				GasMeters:  []GasMeter{},
			},
			true,
		},
		{
			"valid genesis - with incentives",
			&GenesisState{
				Params: DefaultParams(),
				Incentives: []Incentive{
					{
						Contract: "",
						Allocations: sdk.DecCoins{
							sdk.NewDecCoinFromDec("areap", sdk.NewDecWithPrec(5, 2)),
						},
						Epochs:    10,
						StartTime: time.Now(),
					},
				},
			},
			true,
		},
		{
			"invalid genesis - duplicated incentive",
			&GenesisState{
				Params: DefaultParams(),
				Incentives: []Incentive{
					{
						Contract: "",
						Allocations: sdk.DecCoins{
							sdk.NewDecCoinFromDec("areap", sdk.NewDecWithPrec(5, 2)),
						},
						Epochs:    10,
						StartTime: time.Now(),
					},
					{
						Contract: "",
						Allocations: sdk.DecCoins{
							sdk.NewDecCoinFromDec("areap", sdk.NewDecWithPrec(5, 2)),
						},
						Epochs:    10,
						StartTime: time.Now(),
					},
				},
			},
			false,
		},
		{
			"invalid genesis - invalid incentive",
			&GenesisState{
				Params: DefaultParams(),
				Incentives: []Incentive{
					{
						Contract: "",
						Allocations: sdk.DecCoins{
							sdk.NewDecCoinFromDec("areap", sdk.NewDecWithPrec(5, 2)),
						},
						Epochs:    10,
						StartTime: time.Now(),
					},
				},
			},
			false,
		},
		{
			"valid genesis - with gasmeters",
			&GenesisState{
				Params: DefaultParams(),
				GasMeters: []GasMeter{
					{
						Contract:      "",
						Participant:   "",
						CumulativeGas: 10,
					},
				},
			},
			true,
		},
		{
			"invalid genesis - duplicated gasmeter",
			&GenesisState{
				Params: DefaultParams(),
				GasMeters: []GasMeter{
					{
						Contract:      "",
						Participant:   "",
						CumulativeGas: 10,
					},
					{
						Contract:      "",
						Participant:   "",
						CumulativeGas: 10,
					},
				},
			},
			false,
		},
		{
			"invalid genesis - invalid gasmeter",
			&GenesisState{
				Params: DefaultParams(),
				GasMeters: []GasMeter{
					{
						Contract:      "",
						Participant:   "",
						CumulativeGas: 10,
					},
				},
			},
			false,
		},
		{
			"empty genesis",
			&GenesisState{},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}
