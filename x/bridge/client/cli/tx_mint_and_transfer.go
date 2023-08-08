package cli

import (
	"strconv"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"github.com/reapchain/reapchain/v8/x/bridge/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMintAndTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-and-transfer [address] [amount]",
		Short: "Bridge Mint & Transfer Function",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintAndTransfer(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
