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
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/accounts", accountsController.Index).Methods("GET")
	api.HandleFunc("/accounts", accountsController.Create).Methods("POST")
	http.ListenAndServe(":"+port, router)
}
