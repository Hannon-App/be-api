package rents

import "time"

type RentCore struct {
	ID            uint
	StartDate     time.Time
	EndDate       time.Time
	Status        string
	TotalPrice    uint
	Discount      uint
	PaymentLink   *string
	InvoiceNumber string
	IDXendit      string
	UserID        uint
}

type RentDetailCore struct {
	ID       uint
	RentID   uint
	Price    uint
	Quantity uint
}

type RentDataInterface interface {
	Create(input RentCore) error
	GetById(id uint) (RentCore, error)
	AcceptPayment(id uint, userid uint) error
	UpdatebyId(id uint, input RentCore) error
	Callback(input RentCore) error
}

type RentServiceInterface interface {
	Add(input RentCore) error
	ReadById(id uint) (RentCore, error)
	AcceptPayment(id uint, userid uint) error
	UpdatebyId(id uint, input RentCore) error
	Callback(input RentCore) error
}
