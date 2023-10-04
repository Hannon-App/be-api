package data

import (
	"Hannon-app/features/users"
	"Hannon-app/helpers"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

// Insert implements users.UserDataInterface
func (repo *UserQuery) Insert(input users.UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	var userModel = UserCoreToModel(input)

	hash, errHass := helpers.HashPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hash

	if filenameImages == "default.png" {
		userModel.ProfilePhoto = filenameImages
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		userModel.ProfilePhoto = nameGen + filenameImages
		errUp := helpers.Uploader.UploadFile(fileImages, userModel.ProfilePhoto)

		if errUp != nil {
			return errUp
		}
	}

	if filenameID == "default.png" {
		userModel.UploadKTP = filenameID
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		userModel.UploadKTP = nameGen + filenameID
		errUp := helpers.Uploader.UploadFile(fileID, userModel.UploadKTP)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
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

	check := helpers.CheckPassword(password, data.Password)
	if !check {
		return users.UserCore{}, errors.New("password incorect")
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToUserCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}
