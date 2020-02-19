package main

import (
	"net/http"
	"os"
	"transfer-banking/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	p := controller.Accounts{}
	router := http.NewServeMux()

	router.HandleFunc("/", p.Teste)
	http.ListenAndServe(":"+port, router)
}
