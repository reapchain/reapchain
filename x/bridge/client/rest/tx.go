package rest

import (
	"github.com/gorilla/mux"
	"github.com/reapchain/cosmos-sdk/client/tx"
	"net/http"

	"github.com/reapchain/cosmos-sdk/client"
	sdkclientrest "github.com/reapchain/cosmos-sdk/client/rest"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/rest"
	"github.com/reapchain/reapchain/v8/x/bridge/types"
)

// SendReq defines the properties of a send request's body.
type SendReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Address string       `json:"address" yaml:"address"`
	Amount  string       `json:"amount" yaml:"amount"`
}

// NewSendRequestHandlerFn returns an HTTP REST handler for creating a MsgSend
// transaction.
func NewMintAndTransferRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SendReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		argAddress, err := sdk.AccAddressFromBech32(req.Address)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		argAmount, coinParseError := sdk.ParseCoinNormalized(req.Amount)
		if coinParseError != nil {
			return
		}

		msg := types.NewMsgMintAndTransfer(
			clientCtx.GetFromAddress().String(),
			argAddress.String(),
			argAmount.String(),
		)
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)

	}
}

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := sdkclientrest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/bridge/mint-and-transfer", NewMintAndTransferRequestHandlerFn(clientCtx)).Methods("POST")
}
