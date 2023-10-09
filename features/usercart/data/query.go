package data

import (
	"Hannon-app/features/usercart"
	"errors"

	"gorm.io/gorm"
)

type UserCartQuery struct {
	db *gorm.DB
}

// CreateCart implements usercart.UserCartDataInterface.
func (repo *UserCartQuery) CreateCart(input usercart.UserCartCore) error {
	var userCartModel = FromUserCartCore(input)
	tx := repo.db.Create(&userCartModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil

}

// DeleteCart implements usercart.UserCartDataInterface.
func (*UserCartQuery) DeleteCart(id uint) error {
	panic("unimplemented")
}

// SelectAllCart implements usercart.UserCartDataInterface.
func (*UserCartQuery) SelectAllCart() ([]usercart.UserCartCore, error) {
	panic("unimplemented")
}

// SelectCartById implements usercart.UserCartDataInterface.
func (repo *UserCartQuery) SelectCartById(id uint) (usercart.UserCartCore, error) {
	var data UserCart
	repo.db.Where("id = ?", id).Preload("CartItems").First(&data)
	tx := repo.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return usercart.UserCartCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return usercart.UserCartCore{}, errors.New("data not found")
	}
	resultCore := ToUserCartCore(data)
	return resultCore, nil

}

// UpdateCart implements usercart.UserCartDataInterface.
func (*UserCartQuery) UpdateCart(id uint, input usercart.UserCartCore) error {
	panic("unimplemented")
}

func New(db *gorm.DB) usercart.UserCartDataInterface {
	return &UserCartQuery{
		db: db,
	}
}
