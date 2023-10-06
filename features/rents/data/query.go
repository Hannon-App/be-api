package data

import (
	"Hannon-app/features/rents"
	"errors"

	"gorm.io/gorm"
)

type RentQuery struct {
	db *gorm.DB
}

// Callback implements rents.RentDataInterface.
func (repo *RentQuery) Callback(input rents.RentCore) error {
	var data = RentCoreToModel(input)
	tx := repo.db.Where("invoice_number = ?", data.InvoiceNumber).Updates(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// AcceptPayment implements rents.RentDataInterface.
func (repo *RentQuery) AcceptPayment(id uint, userid uint) error {
	var data = RentCoreToModel(rents.RentCore{})
	tx := repo.db.Model(&Rent{}).Where("id, user_id = ?,?", id, userid).Updates(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

// GetById implements rents.RentDataInterface.
func (repo *RentQuery) GetById(id uint) (rents.RentCore, error) {
	var data Rent
	tx := repo.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return rents.RentCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return rents.RentCore{}, errors.New("data not found")
	}

	resultCore := RentModelToCore(data)
	return resultCore, nil
}

// UpdatebyId implements rents.RentDataInterface.
func (repo *RentQuery) UpdatebyId(id uint, input rents.RentCore) error {
	var data = RentCoreToModel(input)
	tx := repo.db.Model(&Rent{}).Where("id = ?", id).Updates(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

// Create implements rents.RentDataInterface.
func (repo *RentQuery) Create(input rents.RentCore) error {
	var data = RentCoreToModel(input)
	tx := repo.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil

}

func New(db *gorm.DB) rents.RentDataInterface {
	return &RentQuery{
		db: db,
	}
}
