package network

import (
	"fmt"

	tmrand "github.com/reapchain/reapchain-core/libs/rand"
	dbm "github.com/tendermint/tm-db"

	"github.com/reapchain/cosmos-sdk/baseapp"
	servertypes "github.com/reapchain/cosmos-sdk/server/types"
	"github.com/reapchain/cosmos-sdk/simapp"
	"github.com/reapchain/cosmos-sdk/simapp/params"
	storetypes "github.com/reapchain/cosmos-sdk/store/types"

	"github.com/reapchain/ethermint/encoding"
	"github.com/reapchain/ethermint/testutil/network"

	"github.com/reapchain/mercury/app"
)

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig() network.Config {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	cfg := network.DefaultConfig()

	cfg.Codec = encCfg.Marshaler
	cfg.TxConfig = encCfg.TxConfig
	cfg.LegacyAmino = encCfg.Amino
	cfg.InterfaceRegistry = encCfg.InterfaceRegistry
	cfg.AppConstructor = NewAppConstructor(encCfg)
	cfg.GenesisState = app.ModuleBasics.DefaultGenesis(encCfg.Marshaler)

	cfg.ChainID = fmt.Sprintf("evmos_%d-1", tmrand.Int63n(9999999999999)+1)
	return cfg
}

// NewAppConstructor returns a new Evmos AppConstructor
func NewAppConstructor(encodingCfg params.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return app.NewEvmos(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encodingCfg,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}
