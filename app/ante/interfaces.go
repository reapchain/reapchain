package ante

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	sdk "github.com/reapchain/cosmos-sdk/types"
	evmtypes "github.com/reapchain/ethermint/x/evm/types"
)

// EvmKeeper defines the expected keeper interface used on the AnteHandler
type EvmKeeper interface {
	GetParams(ctx sdk.Context) (params evmtypes.Params)
	ChainID() *big.Int
	GetBaseFee(ctx sdk.Context, ethCfg *params.ChainConfig) *big.Int
}
