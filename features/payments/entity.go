package payments

import "time"

type VirtualAccountObjectCore struct {
	ID              string
	ExternalID      string
	OwnerID         string
	MerchantCode    string
	BankCode        string
	AccountNumber   string
	Currency        string
	IsSingleUse     *bool
	IsClosed        *bool
	ExpectedAmount  float64
	SuggestedAmount float64
	ExpirationDate  *time.Time
	Name            string
	Description     string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PaymentDataInterface interface {
	CreateVA(input VirtualAccountObjectCore) error
}

type PaymentServiceInterface interface {
	AddVA(input VirtualAccountObjectCore) error
}
