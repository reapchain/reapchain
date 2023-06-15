package cli

import (
	"context"
	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"github.com/spf13/cobra"
)

func CmdShowAllWhitelistedValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-whitelist",
		Short: "Show all whitelisted validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWhitelistedValidatorsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GetAllWhiteListedValidators(context.Background(), params)
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
