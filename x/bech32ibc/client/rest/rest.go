package rest

import (
	"github.com/gorilla/mux"

	"github.com/reapchain/cosmos-sdk/client"
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers bech32ibc-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
}
