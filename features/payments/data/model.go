package data

import (
	"Hannon-app/features/payments"
	"time"
)

type VirtualAccountObject struct {
	ID              string `gorm:"primaryKey"`
	ExternalID      string
	OwnerID         string
	MerchantCode    string
	BankCode        string
	AccountNumber   string
	Currency        string
	IsSingleUse     *bool `gorm:"default:false"`
	IsClosed        *bool `gorm:"default:false"`
	ExpectedAmount  float64
	SuggestedAmount float64
	ExpirationDate  *time.Time
	Name            string
	Description     string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func VAModelToCore(dataModel VirtualAccountObject) payments.VirtualAccountObjectCore {
	var VACore = payments.VirtualAccountObjectCore{
		ID:              dataModel.ID,
		ExternalID:      dataModel.ExternalID,
		OwnerID:         dataModel.OwnerID,
		MerchantCode:    dataModel.MerchantCode,
		BankCode:        dataModel.BankCode,
		AccountNumber:   dataModel.AccountNumber,
		Currency:        dataModel.Currency,
		IsSingleUse:     dataModel.IsSingleUse,
		IsClosed:        dataModel.IsClosed,
		ExpectedAmount:  dataModel.ExpectedAmount,
		SuggestedAmount: dataModel.SuggestedAmount,
		ExpirationDate:  dataModel.ExpirationDate,
		Name:            dataModel.Name,
		Description:     dataModel.Description,
		Status:          dataModel.Status,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
	return VACore
}

func VACoreToModel(dataCore payments.VirtualAccountObjectCore) VirtualAccountObject {
	var VAModel = VirtualAccountObject{
		ID:              dataCore.ID,
		ExternalID:      dataCore.ExternalID,
		OwnerID:         dataCore.OwnerID,
		MerchantCode:    dataCore.MerchantCode,
		BankCode:        dataCore.BankCode,
		AccountNumber:   dataCore.AccountNumber,
		Currency:        dataCore.Currency,
		IsSingleUse:     dataCore.IsSingleUse,
		IsClosed:        dataCore.IsClosed,
		ExpectedAmount:  dataCore.ExpectedAmount,
		SuggestedAmount: dataCore.SuggestedAmount,
		ExpirationDate:  dataCore.ExpirationDate,
		Name:            dataCore.Name,
		Description:     dataCore.Description,
		Status:          dataCore.Status,
		CreatedAt:       dataCore.CreatedAt,
		UpdatedAt:       dataCore.UpdatedAt,
	}
	return VAModel
}
