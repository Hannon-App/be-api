package data

import (
	"Hannon-app/features/admins"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password string
	Role     string `gorm:"default:admin"`
}

func ModelToCore(dataModel Admin) admins.AdminCore {
	var adminCore = admins.AdminCore{
		ID:       dataModel.ID,
		Name:     dataModel.Name,
		Username: dataModel.Username,
		Email:    dataModel.Email,
		Password: dataModel.Password,
		Role:     dataModel.Role,
	}
	return adminCore
}

func AdminCoreToModel(dataCore admins.AdminCore) Admin {
	var adminModel = Admin{
		Model:    gorm.Model{},
		Name:     dataCore.Name,
		Username: dataCore.Username,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
	}
	return adminModel
}
