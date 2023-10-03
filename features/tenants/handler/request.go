package handler

import (
	"Hannon-app/features/tenants"
)

type TenantRequest struct {
	Name      string `json:"name" form:"name"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	IDcard    string `json:"id_card" form:"id_card"`
	Images    string `json:"images" form:"images"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	OpenTime  string `json:"open_time" form:"open_time"`
	CloseTime string `json:"close_time" form:"close_time"`
}

func TenantRequestToCore(input TenantRequest) tenants.TenantCore {
	var TenantCore = tenants.TenantCore{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Phone:     input.Phone,
		Images:    input.Images,
		Address:   input.Address,
		IDcard:    input.IDcard,
		OpenTime:  input.OpenTime,
		CloseTime: input.CloseTime,
	}
	return TenantCore
}
