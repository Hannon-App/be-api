package data

import (
	"Hannon-app/features/payments"
	"errors"
	"log"

	"gorm.io/gorm"
)

type PaymentQuery struct {
	db *gorm.DB
}

// CreateVA implements payments.PaymentDataInterface.
func (repo *PaymentQuery) CreateVA(input payments.VirtualAccountObjectCore) error {
	var data = VACoreToModel(input)
	log.Printf("Attempting to insert: %+v", data)

	tx := repo.db.Create(&data)
	if tx.Error != nil {
		log.Printf("Error occurred: %v", tx.Error)
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		log.Println("No row affected")
		return errors.New("no row affected")
	}

	log.Println("Insert successful")
	return nil
}

func New(db *gorm.DB) payments.PaymentDataInterface {
	return &PaymentQuery{
		db: db,
	}
}
