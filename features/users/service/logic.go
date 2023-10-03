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

// Deletebyid implements users.UserServiceInterface
func (service *UserService) Deletebyid(id uint) error {
	err := service.userData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetUserById implements users.UserServiceInterface
func (service *UserService) GetUserById(id uint) (users.UserCore, error) {
	result, err := service.userData.SelectById(id)
	if err != nil {
		return users.UserCore{}, err
	}
	return result, nil
}

// Add implements users.UserServiceInterface
func (service *UserService) Add(input users.UserCore) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("error validate, user required")
	}

	if len(input.Password) < 8 {
		return errors.New("validation error. password harus minimal 8 characters")
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
		validate: validator.New(),
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
