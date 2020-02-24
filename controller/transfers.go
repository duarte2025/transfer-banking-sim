package controller

import (
	"net/http"
	//"encoding/json"

	//"github.com/gorilla/mux"
	"github.com/gobuffalo/pop"
	"transfer-banking/common"
	"transfer-banking/models"
)

type TransfersController struct {
	common.Controller
}

func (pc *TransfersController) Index(w http.ResponseWriter, r *http.Request) {
	p := pc.GetPaginationParams(r)

	tx := models.DB
	query := pop.Q(tx)

	transfer := models.Transfers{}
	err := query.Paginate(p.Page, p.PerPage).
		All(&transfer)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	count, err := query.Count(models.Transfer{})
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	p.CurrentCount = len(transfer)
	p.TotalCount = count
	pc.SendJSONWithPagination(w, transfer, "transfers", p, 200)
}
