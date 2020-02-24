package controller

import (
	"net/http"
	"github.com/gobuffalo/pop"
	"transfer-banking/common"
	"transfer-banking/models"
)

type AccountsController struct {
	common.Controller
}

func (pc *AccountsController) Index(w http.ResponseWriter, r *http.Request) {
	p := pc.GetPaginationParams(r)

	tx := models.DB
	query := pop.Q(tx)

	users := models.Accounts{}
	err := query.Paginate(p.Page, p.PerPage).
		All(&users)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	count, err := query.Count(models.Accounts{})
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	p.CurrentCount = len(users)
	p.TotalCount = count
	pc.SendJSONWithPagination(w, users, "accounts", p, 200)
}