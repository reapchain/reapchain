package rest

import (
	"net/http"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/tx"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/rest"
	govrest "github.com/reapchain/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/reapchain/cosmos-sdk/x/gov/types"

	"github.com/reapchain/reapchain/v8/x/escrow/types"
)

type RegisterEscrowDenomProposalRequest struct {
	BaseReq            rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title              string       `json:"title" yaml:"title"`
	Description        string       `json:"description" yaml:"description"`
	Deposit            sdk.Coins    `json:"deposit" yaml:"deposit"`
	Denom              string       `json:"denom" yaml:"denom"`
	initialPoolBalance string       `json:"initial_pool_balance" yaml:"initial_supply"`
}

type RegisterEscrowDenomAndConvertProposalRequest struct {
	BaseReq            rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title              string       `json:"title" yaml:"title"`
	Description        string       `json:"description" yaml:"description"`
	Deposit            sdk.Coins    `json:"deposit" yaml:"deposit"`
	Denom              string       `json:"denom" yaml:"denom"`
	initialPoolBalance string       `json:"initial_pool_balance" yaml:"initial_supply"`
}

type ToggleEscrowConversionProposalRequest struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Denom       string       `json:"token" yaml:"token"`
}

type AddToEscrowPoolProposalRequest struct {
	BaseReq     rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title       string       `json:"title" yaml:"title"`
	Description string       `json:"description" yaml:"description"`
	Deposit     sdk.Coins    `json:"deposit" yaml:"deposit"`
	Denom       string       `json:"token" yaml:"token"`
}

func RegisterEscrowDenomProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterEscrowDenomProposalHandler(clientCtx),
	}
}

func RegisterEscrowDenomAndConvertProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newRegisterEscrowDenomAndConvertProposalHandler(clientCtx),
	}
}

func ToggleEscrowConversionRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newToggleEscrowConversionHandler(clientCtx),
	}
}

func AddToEscrowPoolRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: types.ModuleName,
		Handler:  newAddToEscrowPoolProposalHandler(clientCtx),
	}
}

func newRegisterEscrowDenomProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterEscrowDenomProposalRequest

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

		initialPoolAmount, _ := sdk.ParseUint(req.initialPoolBalance)
		content := types.NewRegisterEscrowDenomProposal(req.Title, req.Description, req.Denom, sdk.Int(initialPoolAmount))
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

func newRegisterEscrowDenomAndConvertProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterEscrowDenomAndConvertProposalRequest

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

		initialPoolAmount, _ := sdk.ParseUint(req.initialPoolBalance)
		content := types.NewRegisterEscrowDenoAndConvertmProposal(req.Title, req.Description, req.Denom, sdk.Int(initialPoolAmount), req.BaseReq.From)
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

func newToggleEscrowConversionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ToggleEscrowConversionProposalRequest

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

		content := types.NewToggleEscrowConversionProposal(req.Title, req.Description, req.Denom)
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

func newAddToEscrowPoolProposalHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddToEscrowPoolProposalRequest

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

		content := types.NewAddToEscrowPoolProposal(req.Title, req.Description, req.Denom)
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
