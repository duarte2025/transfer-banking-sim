package controller

import (
	"net/http"
	"transfer-banking/common"
)

type Accounts struct {
	common.Controller
}

func (pc *Accounts) Teste(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Testando!</h1>"))
}
