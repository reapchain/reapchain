package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"github.com/reapchain/mercury/x/reap/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(
		NewMintCoinsTxCmd(),
		NewBurnCoinsTxCmd(),
		NewTransferMintedCoinsTxCmd(),
		NewBurnCoinsFromAccountTxCmd(),
	)
	return cmd 
}

func NewTransferMintedCoinsTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer minted coins [amount] [account]",
		Short: `Transfer minted tokens to Accounts`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			coins := string(args[0])
			account := string(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgTransferMintedCoins(clientCtx.GetFromAddress().String(), coins, account)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewMintCoinsTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mint [amount]",
		Short: `Mint New Tokens`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			coins := string(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgMintCoins(clientCtx.GetFromAddress().String(), coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBurnCoinsTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "burn [amount]",
		Short: `Burn Tokens`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			coins := string(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgBurnCoins(clientCtx.GetFromAddress().String(), coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBurnCoinsFromAccountTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "burn_from_account [amount] [account]",
		Short: `Burn Tokens Directly from a User Account`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			coins := string(args[0])
			account := string(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgBurnCoinsFromAccount(clientCtx.GetFromAddress().String(), coins, account)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
