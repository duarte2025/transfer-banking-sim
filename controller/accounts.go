package controller

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
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

	accounts := models.Accounts{}
	err := query.Paginate(p.Page, p.PerPage).
		All(&accounts)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	count, err := query.Count(models.Accounts{})
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	p.CurrentCount = len(accounts)
	p.TotalCount = count
	pc.SendJSONWithPagination(w, accounts, "accounts", p, 200)
}

type FormAccount struct {
	ID       int64
	Name     string   `json:"name"`
	CPF string   `json:"cpf"`
	Ballance float64   `json:"ballance"`
}

func (pc *AccountsController) Create(w http.ResponseWriter, r *http.Request) {
	var formAccount FormAccount
	err := json.NewDecoder(r.Body).Decode(&formAccount)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	account := models.Account{
		Name:     formAccount.Name,
		CPF: formAccount.CPF,
		Ballance: formAccount.Ballance,
	}

	tx := models.DB
	verrs, err := tx.ValidateAndCreate(&account)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	if verrs.HasAny() {
		pc.SendJSONValidationError(w, verrs)
		return
	}

	pc.SendJSON(w, account, 200)
}

func (pc *AccountsController) Get(w http.ResponseWriter, r *http.Request) {
	p, err := pc.findByID(r)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}
	pc.SendJSON(w, p, 200)
}

func (pc *AccountsController) findByID(r *http.Request) (*models.Account, error) {
	params := mux.Vars(r)
	id := params["account_id"]

	tx := models.DB
	query := pop.Q(tx)

	account :=  models.Account{}
	err := query.Find(&account, id)
	if err != nil {
		return nil, err
	}

	return &account, nil
}