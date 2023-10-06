package handler

import (
	"Hannon-app/features/rents"
)

type RentRequest struct {
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	TotalPrice    uint    `json:"total_price"`
	Discount      uint    `json:"discount"`
	Status        string  `json:"status"`
	PaymentLink   *string `json:"payment_link"`
	InvoiceNumber string  `json:"invoice_number"`
	IDXendit      string  `json:"id_xendit"`
	UserID        uint    `json:"user_id"`
}

func RentRequestToCore(input RentRequest) rents.RentCore {
	var RentCore = rents.RentCore{

		Status:        input.Status,
		TotalPrice:    input.TotalPrice,
		Discount:      input.Discount,
		PaymentLink:   input.PaymentLink,
		InvoiceNumber: input.InvoiceNumber,
		IDXendit:      input.IDXendit,
		UserID:        input.UserID,
	}
	return RentCore
}

type CallBackRequest struct {
	InvoiceNumber string `json:"external_id"`
	Status        string `json:"status"`
}
