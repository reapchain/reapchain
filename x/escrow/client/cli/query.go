package cli

import (
	"context"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

// GetQueryCmd returns the parent command for all escrow CLI query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the escrow module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetRegisteredDenomsCmd(),
		GetEscrowPoolBalanceCmd(),
		GetParamsCmd(),
	)
	return cmd
}

func GetRegisteredDenomsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registered-denoms",
		Short: "Get registered denoms",
		Long:  "Get registered denoms",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryRegisteredDenomsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RegisteredDenoms(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetEscrowPoolBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "escrow-pool-balance [denom]",
		Short: "Get the pool-balance for a registered denomination",
		Long:  "Get the pool-balance for a registered denomination",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if err != nil {
				return err
			}

			req := &types.QueryEscrowPoolBalanceRequest{
				Denom: args[0],
			}

			res, err := queryClient.EscrowPoolBalance(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetParamsCmd queries erc20 module params
func GetParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Gets escrow params",
		Long:  "Gets escrow params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
