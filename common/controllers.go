package common

import (
	"encoding/json"
	"net/http"

	"github.com/gobuffalo/pop"
	"github.com/sirupsen/logrus"
	"github.com/gobuffalo/validate"
)

type Controller struct {
}

type Pagination struct {
	TotalCount   int `json:"total_count"`
	CurrentCount int `json:"count"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
}

func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
	if v == nil {
		c.SendJSON(w, struct{}{}, code)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		logrus.Printf("Error while encoding JSON: %v", err)
		c.SendJSONError(w, err)
		return
	}
}

func (c *Controller) SendJSONWithPagination(w http.ResponseWriter, v interface{}, resourceName string, p *Pagination, code int) {
	body := make(map[string]interface{})

	body["page"] = p
	body[resourceName] = v

	c.SendJSON(w, body, code)
}

func (c *Controller) SendJSONError(w http.ResponseWriter, err error) {
	c.SendJSONErrorWithCode(w, err, http.StatusBadRequest)
}

func (c *Controller) SendJSONErrorWithCode(w http.ResponseWriter, err error, code int) {
	var e = struct {
		Error string `json:"error"`
	}{
		err.Error(),
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}

func (c *Controller) GetPaginationParams(r *http.Request) *Pagination {
	paginator := pop.NewPaginatorFromParams(r.URL.Query())

	return &Pagination{
		Page:    paginator.Page,
		PerPage: paginator.PerPage,
	}
}

func (c *Controller) SendJSONValidationError(w http.ResponseWriter, verrs *validate.Errors) {
	var e = struct {
		Error  string              `json:"error"`
		Errors map[string][]string `json:"errors"`
	}{
		"Validation errors",
		verrs.Errors,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(e)
}

