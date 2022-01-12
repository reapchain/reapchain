package keeper

import (
	"fmt"

	"github.com/reapchain/reapchain-core/libs/log"

	"github.com/reapchain/cosmos-sdk/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	paramtypes "github.com/reapchain/cosmos-sdk/x/params/types"
	"github.com/reapchain/mercury/x/reap/types"
	
)

type (
	Keeper struct {
		
		cdc      	codec.BinaryCodec
		storeKey 	sdk.StoreKey
		memKey   	sdk.StoreKey
		paramstore	paramtypes.Subspace
		bankKeeper       types.BankKeeper
		accountKeeper       types.AccountKeeper
		
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	authKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
    
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		
		cdc:      	cdc,
		storeKey: 	storeKey,
		memKey:   	memKey,
		paramstore:	ps,
		accountKeeper: authKeeper,
		bankKeeper: bankKeeper,
		
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintCoins(ctx sdk.Context, msg types.MintCoins) error {

	var mint = types.MintCoins{
		Creator: msg.Creator,
		Amount: msg.Amount,
	}

	newCoins, err := sdk.ParseCoinsNormalized(mint.Amount)
	if err != nil {
		panic(err)
	}

	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}




func (k Keeper) TransferMintedCoins(ctx sdk.Context, msg types.TransferMintedCoins) error {

	var transferMinted = types.TransferMintedCoins{
		Creator: msg.Creator,
		Amount: msg.Amount,
		Account: msg.Account,
	}

	transferCoins, err := sdk.ParseCoinsNormalized(transferMinted.Amount)
	if err != nil {
		panic(err)
	}
	if transferCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	recipient, err := sdk.AccAddressFromBech32(transferMinted.Account)
	if err != nil {
		panic(err)
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, transferCoins)
}

func (k Keeper) BurnCoins(ctx sdk.Context, msg types.BurnCoins) error {

	var burn = types.BurnCoins{
		Creator: msg.Creator,
		Amount: msg.Amount,
	}

	burnCoin, err := sdk.ParseCoinsNormalized(burn.Amount)
	if err != nil {
		panic(err)
	}

	if burnCoin.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoin)
}

func (k Keeper) BurnCoinsFromAccount(ctx sdk.Context, msg types.BurnCoinsFromAccount) error {

	var burnFromAccount = types.BurnCoinsFromAccount{
		Creator: msg.Creator,
		Amount: msg.Amount,
		Account: msg.Account,
	}

	burnAmount, err := sdk.ParseCoinsNormalized(burnFromAccount.Amount)
	if err != nil {
		panic(err)
	}

	if burnAmount.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	sender, err := sdk.AccAddressFromBech32(burnFromAccount.Account)
	if err != nil {
		panic(err)
	}


	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnAmount)
	if transferErr != nil {
		panic(err)
	}

	return k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnAmount)
}
