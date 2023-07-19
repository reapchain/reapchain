package cli

import (
	"fmt"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/tx"
	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/version"
	"github.com/reapchain/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"github.com/spf13/cobra"
)

// NewRegisterCoinProposalCmd implements the command to submit a community-pool-spend proposal
func CmdRegisterStandingMemberProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-standing-member [validator-address] [account-address] [moniker-address]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit a register-standing-member proposal",
		Long: `Submit a proposal to register a Validator to the White List.
Upon passing, the validator-addresss submitted will be allowed to become a full standing member validator`,
		Example: fmt.Sprintf(`$ %s tx gov submit-proposal register-standing-member  <validator-operating-address> --from=<key_or_address>
}`, version.AppName,
		),
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

			validatorAddresss, err := sdk.ValAddressFromBech32(args[0])

			if err != nil {
				return err
			}

			accountAddress, err := sdk.AccAddressFromBech32(args[1])

			if err != nil {
				return err
			}
			moniker := args[2]
			if len(moniker) > stakingtypes.MaxMonikerLength {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid moniker length; got: %d, max: %d", len(moniker), stakingtypes.MaxMonikerLength)
			}
			if len(moniker) < 5 {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid moniker length; got: %d, minimum: %d", len(moniker), 5)
			}

			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content := types.NewMsgRegisterStandingMemberProposal(title, description, validatorAddresss.String(), accountAddress.String(), moniker)

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
