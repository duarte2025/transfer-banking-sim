package controller

import (
	"net/http"
	"encoding/json"
	"fmt"
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

type FormTransfer struct {
	OriginId int64 `json:"account_origin_id"`
	DestinationId int64 `json:"account_destination_id"`
	Amount float64 `json:"amount"`
}

func (pc *TransfersController) Create(w http.ResponseWriter, r *http.Request) {
	var formTransfer FormTransfer
	err := json.NewDecoder(r.Body).Decode(&formTransfer)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	tx := models.DB
	query1 := pop.Q(tx)
	query2 := pop.Q(tx)

	origin := models.Account{}
	destination := models.Account{}

	err = query1.Find(&origin, formTransfer.OriginId)
	if err != nil {
		return
	}

	if origin.Ballance < formTransfer.Amount {
		pc.SendJSON(w, "Saldo insufiente!", 402)
		return 
	}

	err = query2.Find(&destination, formTransfer.DestinationId)
	if err != nil {
		fmt.Printf("%f", destination.Ballance);
		return
	}

	origin.Ballance = origin.Ballance - formTransfer.Amount
	destination.Ballance = destination.Ballance + formTransfer.Amount
	transfer := models.Transfer{
		OriginId:     formTransfer.OriginId,
		DestinationId: formTransfer.DestinationId,
		Amount: formTransfer.Amount,
	}

	verrs, err := tx.ValidateAndCreate(&transfer)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	verrs, err = tx.ValidateAndUpdate(&origin)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}


	verrs, err = tx.ValidateAndUpdate(&destination)
	if err != nil {
		pc.SendJSONError(w, err)
		return
	}

	if verrs.HasAny() {
		pc.SendJSONValidationError(w, verrs)
		return
	}

	
	pc.SendJSON(w, transfer, 200)
}
