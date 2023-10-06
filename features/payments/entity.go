package payments

type PaymentCore struct {
	ID           uint
	CheckoutLink string
	ExternalID   uint
	Status       string
}
