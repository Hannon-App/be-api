package data

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	CheckoutLink string
	ExternalID   uint
	Status       string
}
