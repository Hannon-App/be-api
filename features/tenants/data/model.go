package data

import (
	"Hannon-app/features/tenants"
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Phone     string
	Role      string `gorm:"default:tenant"`
	Address   string
	Images    string
	IDcard    string
	OpenTime  string
	CloseTime string
}

func TenantModelToCore(dataModel Tenant) tenants.TenantCore {
	var TenantCore = tenants.TenantCore{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Phone:     dataModel.Phone,
		Images:    dataModel.Images,
		Address:   dataModel.Address,
		IDcard:    dataModel.IDcard,
		OpenTime:  dataModel.OpenTime,
		CloseTime: dataModel.CloseTime,
		Role:      dataModel.Role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
	return TenantCore
}

func TenantCoreToModel(dataCore tenants.TenantCore) Tenant {
	var TenantModel = Tenant{
		Model:     gorm.Model{},
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Role:      dataCore.Role,
		Phone:     dataCore.Phone,
		Address:   dataCore.Address,
		Images:    dataCore.Images,
		IDcard:    dataCore.IDcard,
		OpenTime:  dataCore.OpenTime,
		CloseTime: dataCore.CloseTime,
	}
	return TenantModel
}
