package ante

import (
	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/types/tx/signing"
	"github.com/reapchain/cosmos-sdk/x/auth/ante"
	authsigning "github.com/reapchain/cosmos-sdk/x/auth/signing"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"
	govkeeper "github.com/reapchain/cosmos-sdk/x/gov/keeper"
	ibcante "github.com/reapchain/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/reapchain/ibc-go/v3/modules/core/keeper"

	ethante "github.com/reapchain/ethermint/app/ante"
	evmtypes "github.com/reapchain/ethermint/x/evm/types"

	permissionskeeper "github.com/reapchain/reapchain/v8/x/permissions/keeper"
	vestingtypes "github.com/reapchain/reapchain/v8/x/vesting/types"
)

// HandlerOptions defines the list of module keepers required to run the Reapchain
// AnteHandler decorators.
type HandlerOptions struct {
	GovKeeper         govkeeper.Keeper
	PermissionsKeeper permissionskeeper.Keeper
	AccountKeeper     evmtypes.AccountKeeper
	BankKeeper        evmtypes.BankKeeper
	IBCKeeper         *ibckeeper.Keeper
	FeeMarketKeeper   evmtypes.FeeMarketKeeper
	StakingKeeper     vestingtypes.StakingKeeper
	EvmKeeper         ethante.EVMKeeper
	FeegrantKeeper    ante.FeegrantKeeper
	SignModeHandler   authsigning.SignModeHandler
	SigGasConsumer    func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	Cdc               codec.BinaryCodec
	MaxTxGasWanted    uint64
}

// Validate checks if the keepers are defined
func (options HandlerOptions) Validate() error {
	if options.AccountKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.StakingKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "staking keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.FeeMarketKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "fee market keeper is required for AnteHandler")
	}
	if options.EvmKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "evm keeper is required for AnteHandler")
	}
	return nil
}

// newCosmosAnteHandler creates the default ante handler for Ethereum transactions
func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.NewEthSetUpContextDecorator(options.EvmKeeper),                         // outermost AnteDecorator. SetUpContext must be called first
		ethante.NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against the node's minimal-gas-prices config
		ethante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice
		ethante.NewEthValidateBasicDecorator(options.EvmKeeper),
		ethante.NewEthSigVerificationDecorator(options.EvmKeeper),
		ethante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		ethante.NewCanTransferDecorator(options.EvmKeeper),
		NewEthVestingTransactionDecorator(options.AccountKeeper),
		ethante.NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
		ethante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		ethante.NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

// newCosmosAnteHandler creates the default ante handler for Cosmos transactions
func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewMempoolFeeDecorator(),
		ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		NewVestingDelegationDecorator(options.AccountKeeper, options.BankKeeper, options.StakingKeeper, options.Cdc),
		NewValidatorCommissionDecorator(options.Cdc),

		NewCreateValidatorMessage(options.StakingKeeper, options.Cdc, options.PermissionsKeeper),
		NewGovernanceSubmitProposalMessage(options.Cdc, options.PermissionsKeeper, options.GovKeeper),
		NewStakingDelegationMessage(options.StakingKeeper, options.Cdc, options.PermissionsKeeper),

		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}

// newCosmosAnteHandlerEip712 creates the ante handler for transactions signed with EIP712
func newCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ethante.NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		NewVestingDelegationDecorator(options.AccountKeeper, options.BankKeeper, options.StakingKeeper, options.Cdc),
		NewValidatorCommissionDecorator(options.Cdc),

		NewCreateValidatorMessage(options.StakingKeeper, options.Cdc, options.PermissionsKeeper),
		NewGovernanceSubmitProposalMessage(options.Cdc, options.PermissionsKeeper, options.GovKeeper),
		NewStakingDelegationMessage(options.StakingKeeper, options.Cdc, options.PermissionsKeeper),

		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		ethante.NewEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}
