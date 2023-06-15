package rest

import (
	"net/http"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/tx"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/rest"
	govrest "github.com/reapchain/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	"github.com/reapchain/reapchain/v8/x/permissions/types"
)

// RegisterStandingMemberProposalRequest defines a request for a new register coin proposal.
type RegisterStandingMemberProposalRequest struct {
	BaseReq          rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title            string       `json:"title" yaml:"title"`
	Description      string       `json:"description" yaml:"description"`
	AccountAddress   string       `json:"account_address" yaml:"account_address"`
	ValidatorAddress string       `json:"validator_address" yaml:"validator_address"`
	Moniker          string       `json:"moniker" yaml:"moniker"`
	Deposit          sdk.Coins    `json:"deposit" yaml:"deposit"`
}

// RemoveStandingMemberProposalRequest defines a request for a new register coin proposal.
type RemoveStandingMemberProposalRequest struct {
	BaseReq          rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title            string       `json:"title" yaml:"title"`
	Description      string       `json:"description" yaml:"description"`
	ValidatorAddress string       `json:"sender" yaml:"sender"`
	Deposit          sdk.Coins    `json:"deposit" yaml:"deposit"`
}

// RegisterStandingMemberProposalRequest defines a request for a new register coin proposal.
type ReplaceStandingMemberProposalRequest struct {
	BaseReq                     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title                       string       `json:"title" yaml:"title"`
	Description                 string       `json:"description" yaml:"description"`
	ExistingValidatorAddress    string       `json:"existing_validator_address" yaml:"existing_validator_address"`
	ReplacementValidatorAddress string       `json:"replacement_validator_address" yaml:"replacement_validator_address"`
	ReplacementAccountAddress   string       `json:"replacement_account_address" yaml:"replacement_account_address"`
	Moniker                     string       `json:"moniker" yaml:"moniker"`
	Deposit                     sdk.Coins    `json:"deposit" yaml:"deposit"`
}

func RegisterStandingMemberProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterStandingMemberProposalHandler(clientCtx),
	}
}

func RemoveStandingMemberProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRemoveStandingMemberProposalHandler(clientCtx),
	}
}

func ReplaceStandingMemberProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newReplaceStandingMemberProposalHandler(clientCtx),
	}
}

func newRegisterStandingMemberProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterStandingMemberProposalRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewMsgRegisterStandingMemberProposal(req.Title, req.Description, req.ValidatorAddress, req.AccountAddress, req.Moniker)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newRemoveStandingMemberProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RemoveStandingMemberProposalRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewMsgRemoveStandingMemberProposal(req.Title, req.Description, req.ValidatorAddress)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func newReplaceStandingMemberProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ReplaceStandingMemberProposalRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.NewMsgReplaceStandingMemberProposal(req.Title, req.Description, req.ExistingValidatorAddress, req.ReplacementValidatorAddress, req.ReplacementAccountAddress, req.Moniker)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
