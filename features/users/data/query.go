package data

import (
	"Hannon-app/features/users"
	"errors"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

// Insert implements users.UserDataInterface
func (repo *UserQuery) Insert(input users.UserCore) error {
	inputModel := UserCoreToModel(input)

	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

func (repo *UserQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User

	tx := repo.db.Where("email = ?", email).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToUserCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
