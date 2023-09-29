package data

import (
	"Hannon-app/features/admins"
	"errors"

	"gorm.io/gorm"
)

type AdminQuery struct {
	db        *gorm.DB
	dataLogin admins.AdminCore
}

func New(db *gorm.DB) admins.AdminDataInterface {
	return &AdminQuery{
		db: db,
	}
}

// Login implements admins.AdminDataInterface.
func (repo *AdminQuery) Login(email string, password string) (dataLogin admins.AdminCore, err error) {
	var data Admin
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return admins.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return admins.AdminCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
