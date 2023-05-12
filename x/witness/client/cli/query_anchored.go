package cli

import (
	"context"
	"strconv"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func CmdListAnchored() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-anchored",
		Short: "list all anchored",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAnchoredRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AnchoredAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAnchored() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-anchored [id]",
		Short: "shows a anchored",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAnchoredRequest{
				Id: id,
			}

			res, err := queryClient.Anchored(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
