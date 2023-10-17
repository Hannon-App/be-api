package handler

import "time"

type VAResponse struct {
	ID             string     `json:"id"`
	ExternalID     string     `json:"external_id"`
	OwnerID        string     `json:"owner_id"`
	BankCode       string     `json:"bank_code"`
	MerchantCode   string     `json:"merchant_code"`
	AccountNumber  string     `json:"account_number"`
	Name           string     `json:"name"`
	Currency       string     `json:"currency"`
	Country        string     `json:"country,omitempty"`
	IsSingleUse    *bool      `json:"is_single_use,omitempty"`
	IsClosed       *bool      `json:"is_closed,omitempty"`
	ExpirationDate *time.Time `json:"expiration_date"`
	Status         string     `json:"status"`
}
