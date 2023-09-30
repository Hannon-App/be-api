package service

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/users"
	"errors"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

// Add implements users.UserServiceInterface
func (service *UserService) Add(input users.UserCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, user required")
	}
	if input.Name == "" {
		return errors.New("validation error. name required")
	}
	if input.UserName != "" {
		return errors.New("validation error. username required")
	}
	if input.Email != "" {
		return errors.New("validation error. email required")
	}
	if input.Password != "" {
		return errors.New("validation error. password required")
	}
	if len(input.Password) <= 7 {
		return errors.New("validation error. password harus minimal 8 characters")
	}

	if input.PhoneNumber != "" {
		return errors.New("validation error. Nomor telepon required")
	}
	if input.ProfilePhoto != "" {
		return errors.New("validation error. foto Profil belum diisi")
	}
	if input.UploadKTP != "" {
		return errors.New("validation error. Foto KTP belum diisi")
	}
	errInsert := service.userData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func New(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

func (service *UserService) Login(email string, password string) (dataLogin users.UserCore, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.ID, dataLogin.ID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}
