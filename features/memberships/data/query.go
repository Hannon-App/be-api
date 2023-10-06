package data

import (
	"Hannon-app/features/memberships"
	"errors"

	"gorm.io/gorm"
)

type MembershipQuery struct {
	db        *gorm.DB
	dataLogin memberships.MembershipCore
}

func New(db *gorm.DB) memberships.MembershipDataInterface {
	return &MembershipQuery{
		db: db,
	}
}

// Login implements memberships.MembershipDataInterface.
func (repo *MembershipQuery) Login(email string, password string) (dataLogin memberships.MembershipCore, err error) {
	var data Membership
	tx := repo.db.Where("email = ? and password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return memberships.MembershipCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return memberships.MembershipCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
