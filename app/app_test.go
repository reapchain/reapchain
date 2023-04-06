package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/simapp"

	abci "github.com/reapchain/reapchain-core/abci/types"
	"github.com/reapchain/reapchain-core/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/reapchain/ethermint/encoding"
	"github.com/reapchain/reapchain/v8/types"
)

func TestEvmosExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewEvmos(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encoding.MakeConfig(ModuleBasics), simapp.EmptyAppOptions{})

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			ChainId:       types.MainnetChainID + "-1",
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewEvmos(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, encoding.MakeConfig(ModuleBasics), simapp.EmptyAppOptions{})
	_, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
