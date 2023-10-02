package data

import (
	"Hannon-app/features/users"
	"Hannon-app/helpers"
	"errors"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

// Delete implements users.UserDataInterface
func (repo *UserQuery) Delete(id uint) error {
	var userGorm User
	tx := repo.db.First(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	// Hapus pengguna dari database
	tx = repo.db.Delete(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found to deleted")

	}
	return nil
}

// SelectById implements users.UserDataInterface
func (repo *UserQuery) SelectById(id uint) (users.UserCore, error) {
	var result User
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("data not found")
	}

	resultCore := ModelToUserCore(result)
	return resultCore, nil
}

// Insert implements users.UserDataInterface
func (repo *UserQuery) Insert(input users.UserCore) error {
	inputModel := UserCoreToModel(input)

	hass, errHass := helpers.HassPassword(inputModel.Password)
	if errHass != nil {
		return errHass
	}
	inputModel.Password = hass

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
		db:        db,
		dataLogin: users.UserCore{},
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
