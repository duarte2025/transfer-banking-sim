package main

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"transfer-banking/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := mux.NewRouter()

	accountsController := controller.AccountsController{}
	transfersController := controller.TransfersController{}

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/accounts", accountsController.Index).Methods("GET")
	api.HandleFunc("/accounts/{account_id}/ballance", accountsController.Get).Methods("GET")
	api.HandleFunc("/accounts", accountsController.Create).Methods("POST")

	api.HandleFunc("/transfers", transfersController.Index).Methods("GET")
	api.HandleFunc("/transfers", transfersController.Create).Methods("POST")

	http.ListenAndServe(":"+port, router)
}
