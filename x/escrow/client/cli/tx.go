package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/version"
	"github.com/reapchain/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// NewTxCmd returns a root CLI command handler for erc20 transaction commands
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "escrow subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConvertToNativeCmd(),
		NewConvertToDenomCmd(),
	)
	return txCmd
}

func NewConvertToNativeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "convert-to-native [coin] [receiver]",
		Short:   "Convert a coin with a registered denomination to the Default denomination. When the receiver [optional] is omitted, the coins are transferred to the sender.",
		Example: fmt.Sprintf(`$ %s tx escrow convert-to-native <registered_denomination_and_amount> <optional_reciever_address> --from <key_or_address> `, version.AppName),
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress().String()
			var receiver string
			if len(args) == 2 {
				receiverAddr, err := sdk.AccAddressFromBech32(args[1])
				if err != nil {
					return err
				}
				receiver = receiverAddr.String()

			} else {
				receiver = sender
			}

			msg := &types.MsgConvertToNative{
				Coin:     coin,
				Sender:   sender,
				Receiver: receiver,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewConvertToDenomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "convert-to-denom [amount] [denom] [receiver] ",
		Short:   "Convert a specified amount to a Registered Denomination. When the receiver [optional] is omitted, the coins are transferred to the sender.",
		Args:    cobra.RangeArgs(2, 3),
		Long:    `Convert your balance of native chain denomination to a specified amount of a specific denomination. When the receiver [optional] is omitted, the coins are transferred to the sender.`,
		Example: fmt.Sprintf(`$ %s tx escrow convert-to-denom 1000 <registered_denomination> --from <key_or_address> `, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid amount %s", args[1])
			}

			denom := args[1]

			sender := cliCtx.GetFromAddress().String()
			var receiver string

			if len(args) == 3 {
				receiverAddr, err := sdk.AccAddressFromBech32(args[2])
				if err != nil {
					return err
				}
				receiver = receiverAddr.String()

			} else {
				receiver = sender
			}

			msg := &types.MsgConvertToDenom{
				Amount:   amount,
				Denom:    denom,
				Sender:   sender,
				Receiver: receiver,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRegisterEscrowDenomProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-escrow-denom [denomination] [initial-supply]",
		Args:    cobra.ExactArgs(2),
		Short:   "Submit a proposal to register a denom for Escrow Authorization",
		Long:    "Submit a proposal to register a denom for Escrow Authorization along with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx gov submit-proposal register-escrow-denom  <denomination> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			denomination := args[0]
			initialSupply, _ := sdk.ParseUint(args[1])
			from := clientCtx.GetFromAddress()
			content := types.NewRegisterEscrowDenomProposal(title, description, denomination, sdk.Int(initialSupply))

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1areap", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}

func NewRegisterEscrowDenomAndConvertProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register-escrow-denom-and-convert [denomination] [initial-supply]",
		Args:    cobra.ExactArgs(2),
		Short:   "Submit a proposal to register a denom for Escrow Authorization and automatically convert to specified denom",
		Long:    "Submit a proposal to register a denom for Escrow Authorization and conversion along with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx gov submit-proposal register-escrow-denom  <denomination> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			denomination := args[0]
			initialSupply, _ := sdk.ParseUint(args[1])
			from := clientCtx.GetFromAddress()
			content := types.NewRegisterEscrowDenoAndConvertmProposal(title, description, denomination, sdk.Int(initialSupply), from.String())

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1areap", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}

func NewToggleEscrowConversionProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "toggle-escrow-conversion [denomination]",
		Args:    cobra.ExactArgs(1),
		Short:   "Submit a proposal to toggle conversion of a denomination",
		Long:    "Submit a proposal to toggle the conversion of a denomination  with an initial deposit.",
		Example: fmt.Sprintf("$ %s tx gov submit-proposal toggle-token-conversion <denomination> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			denom := args[0]
			content := types.NewToggleEscrowConversionProposal(title, description, denom)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1areap", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}

func NewAddToEscrowPoolProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-escrow-supply [denomination]",
		Args:    cobra.ExactArgs(1),
		Short:   "Submit a proposal to add circulation supply for an Escrow Registered Denomination",
		Long:    "Submit a proposal to add circulation supply for an Escrow Registered Denomination",
		Example: fmt.Sprintf("$ %s tx gov submit-proposal add-escrow-supply <denomination> --from=<key_or_address>", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			denom := args[0]
			content := types.NewAddToEscrowPoolProposal(title, description, denom)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1areap", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}
	return cmd
}
