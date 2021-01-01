package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers tictactoe-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/tictactoe/games/{id}", getGameHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/tictactoe/games", listGameHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/tictactoe/games", createGameHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/tictactoe/games/{id}", updateGameHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/tictactoe/games/{id}", deleteGameHandler(clientCtx)).Methods("POST")

}

