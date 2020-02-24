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

	p := controller.AccountsController{}
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/accounts", p.Index).Methods("GET")
	http.ListenAndServe(":"+port, router)
}
