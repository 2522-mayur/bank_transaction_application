package router

import (
	"github.com/gorilla/mux"
	"BANK_TRANSACTION_APPLICATION/handler"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", handler.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", handler.CreateTransaction).Methods("POST")
	return r
}