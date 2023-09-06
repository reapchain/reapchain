package cli

import (
	"strconv"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

var _ = strconv.Itoa(0)

func CmdAnchoring() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "anchoring [block-hash] [block-height] [chain-id] [filter]",
		Short: "Broadcast message anchoring",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlockHash := args[0]
			argBlockHeight := args[1]
			argChainID := args[2]
			argFilter := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAnchoring(
				clientCtx.GetFromAddress().String(),
				argBlockHash,
				argBlockHeight,
				argChainID,
				argFilter,
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
