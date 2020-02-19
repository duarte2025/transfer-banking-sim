package common

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Controller struct {
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

func (c *Controller) SendJSONError(w http.ResponseWriter, err error) {
	c.SendJSONErrorWithCode(w, err, http.StatusBadRequest)
}
