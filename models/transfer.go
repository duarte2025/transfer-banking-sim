package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
)
// Transfer is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Transfer struct {
	ID int64 `json:"id" db:"id"`
	OriginId int64 `json:"account_origin_id" db:"account_origin_id"`
	DestinationId int64 `json:"account_destination_id" db:"account_destination_id"`
	Amount float64 `json:"amount" db:"amount"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// String is not required by pop and may be deleted
func (t Transfer) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Transfers is not required by pop and may be deleted
type Transfers []Transfer

// String is not required by pop and may be deleted
func (t Transfers) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Transfer) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Transfer) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Transfer) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
