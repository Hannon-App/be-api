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

// Delete implements users.UserDataInterface
func (repo *UserQuery) Delete(adminID uint, id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// ReadAll implements users.UserDataInterface
func (repo *UserQuery) ReadAll(adminID uint, page uint, userPerPage uint, searchName string) ([]users.UserCore, int64, error) {
	var userData []User
	var totalCount int64

	if page == 0 && userPerPage == 0 {
		tx := repo.db

		if searchName != "" {
			tx = tx.Where("name LIKE ?", "%"+searchName+"%")
		}
		tx.Find(&userData)
	} else {

		offset := int((page - 1) * userPerPage)

		query := repo.db.Offset(offset).Limit(int(userPerPage))

		if searchName != "" {
			query = query.Where("name LIKE ?", "%"+searchName+"%")
		}

		tx := query.Find(&userData)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
	}

	var userCore []users.UserCore
	for _, value := range userData {
		userCore = append(userCore, ModelToUserCore(value))
	}

	repo.db.Model(&User{}).Count(&totalCount)

	return userCore, totalCount, nil
}

// UpdateUser implements users.UserDataInterface
func (repo *UserQuery) UpdateUser(uID uint, id uint, input users.UserCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	var user User
	tx := repo.db.Where("id = ? AND id = ?", id, uID).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	// Mapping Entity Target to Model
	updatedUser := UserCoreToModel(input)

	// Hash the password
	hashedPassword, err := helpers.HashPassword(updatedUser.Password)
	if err != nil {
		return err
	}
	updatedUser.Password = hashedPassword

	// Check if profil photo file is provided
	if filenameImages != "default.png" {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		updatedUser.ProfilePhoto = nameGen + filenameImages

		// Upload and replace the profil photo
		errUp := helpers.Uploader.UploadFile(fileImages, updatedUser.ProfilePhoto)
		if errUp != nil {
			return errUp
		}
	}

	// Check if KTP photo file is provided
	if filenameID != "default.png" {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		updatedUser.UploadKTP = nameGen + filenameID

		// Upload and replace the KTP photo
		errUp := helpers.Uploader.UploadFile(fileID, updatedUser.UploadKTP)
		if errUp != nil {
			return errUp
		}
	}

	// Perform the update of user data in the database
	tx = repo.db.Model(&user).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
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
