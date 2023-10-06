package data

import (
	"Hannon-app/features/rents"
	"time"

	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
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

func RentModelToCore(dataModel Rent) rents.RentCore {
	var rentCore = rents.RentCore{
		ID:            dataModel.ID,
		StartDate:     dataModel.StartDate,
		EndDate:       dataModel.EndDate,
		Status:        dataModel.Status,
		TotalPrice:    dataModel.TotalPrice,
		Discount:      dataModel.Discount,
		PaymentLink:   dataModel.PaymentLink,
		InvoiceNumber: dataModel.InvoiceNumber,
		IDXendit:      dataModel.IDXendit,
		UserID:        dataModel.UserID,
	}
	return rentCore
}

func RentCoreToModel(dataCore rents.RentCore) Rent {
	var rentModel = Rent{
		Model:         gorm.Model{},
		StartDate:     dataCore.StartDate,
		EndDate:       dataCore.EndDate,
		Status:        dataCore.Status,
		TotalPrice:    dataCore.TotalPrice,
		Discount:      dataCore.Discount,
		PaymentLink:   dataCore.PaymentLink,
		InvoiceNumber: dataCore.InvoiceNumber,
		IDXendit:      dataCore.IDXendit,
		UserID:        dataCore.UserID,
	}
	return rentModel
}
