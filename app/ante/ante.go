package ante

import (
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	authante "github.com/reapchain/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/reapchain/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/reapchain/cosmos-sdk/x/bank/keeper"
	feegrantkeeper "github.com/reapchain/cosmos-sdk/x/feegrant/keeper"
	ibcante "github.com/reapchain/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/reapchain/ibc-go/v3/modules/core/keeper"

	ethante "github.com/reapchain/ethermint/app/ante"

	gravitykeeper "github.com/reapchain/reapchain/v8/x/gravity/keeper"
)

// NewAnteHandler returns an ante handler responsible for attempting to route an
// Ethereum or SDK transaction to an internal ante handler for performing
// transaction-level processing (e.g. fee payment, signature verification) before
// being passed onto it's respective handler.
func NewAnteHandler(
	options HandlerOptions,
	gravityKeeper *gravitykeeper.Keeper,
	accountKeeper *authkeeper.AccountKeeper,
	bankKeeper *bankkeeper.BaseKeeper,
	feegrantKeeper *feegrantkeeper.Keeper,
	ibcKeeper *ibckeeper.Keeper,
	cdc codec.BinaryCodec,
) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		defer ethante.Recover(ctx.Logger(), &err)

		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx
					anteHandler = newEthAnteHandler(options)
				case "/ethermint.types.v1.ExtensionOptionsWeb3Tx":
					// handle as normal Cosmos SDK tx, except signature is checked for EIP712 representation
					anteHandler = newCosmosAnteHandlerEip712(options)
				default:
					return ctx, sdkerrors.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s", typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}

		// handle as totally normal Cosmos SDK tx
		switch tx.(type) {
		case sdk.Tx:
			anteHandler = newCosmosAnteHandler(options)
		default:
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		// Create additional AnteDecorators to chain together
		ibcAnteDecorator := ibcante.NewAnteDecorator(ibcKeeper)
		minCommissionDecorator := NewMinCommissionDecorator(cdc)

		addlDecorators := []sdk.AnteDecorator{ibcAnteDecorator, minCommissionDecorator}
		// Chain together and terminate the input decorators array
		customHandler := sdk.ChainAnteDecorators(addlDecorators...)

		// Create and return a function which ties the two handlers together
		fullHandler := chainHandlers(anteHandler, customHandler)

		return fullHandler(ctx, tx, sim)
	}
}

// Loosely chain together two AnteHandlers by first calling handler, then calling secondHandler with the result
// NOTE: This order is important due to the way GasMeter works, see sdk.ChainAnteDecorators for more info
// This process is necessary because AnteHandlers are immutable once constructed, and the sdk does not expose its
// curated list of default AnteDecorators.
func chainHandlers(handler sdk.AnteHandler, secondHandler sdk.AnteHandler) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
		// First run the original AnteHandler (likely the default sdk chain of decorators)
		newCtx, err := handler(ctx, tx, simulate)
		if err != nil { // Return early from an error
			return newCtx, err
		}

		// Next run the second handler
		return secondHandler(newCtx, tx, simulate)
	}
}
