package handler

import "time"

type VARequest struct {
	ExternalID           string     `json:"external_id" validate:"required"`
	BankCode             string     `json:"bank_code" validate:"required"`
	Name                 string     `json:"name" validate:"required"`
	VirtualAccountNumber string     `json:"virtual_account_number,omitempty"`
	IsClosed             *bool      `json:"is_closed,omitempty"`
	IsSingleUse          *bool      `json:"is_single_use,omitempty"`
	ExpirationDate       *time.Time `json:"expiration_date,omitempty"`
	SuggestedAmount      float64    `json:"suggested_amount,omitempty"`
	ExpectedAmount       float64    `json:"expected_amount,omitempty"`
	Description          string     `json:"description,omitempty"`
}
