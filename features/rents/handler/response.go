package handler

type RentResponse struct {
	ID            uint    `json:"id"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	Status        string  `json:"status"`
	TotalPrice    uint    `json:"total_price"`
	Discount      uint    `json:"discount"`
	PaymentLink   *string `json:"payment_link"`
	InvoiceNumber string  `json:"invoice_number"`
	UserID        uint    `json:"user_id"`
}
