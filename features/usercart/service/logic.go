package service

import (
	"Hannon-app/features/usercart"

	"github.com/go-playground/validator/v10"
)

type UserCartService struct {
	UserCartData usercart.UserCartDataInterface
	validate     *validator.Validate
}

// CreateCart implements usercart.UserCartServiceInterface.
func (service *UserCartService) CreateCart(input usercart.UserCartCore) error {
	err := service.UserCartData.CreateCart(input)
	return err
}

// DeleteCart implements usercart.UserCartServiceInterface.
func (*UserCartService) DeleteCart(id uint) error {
	panic("unimplemented")
}

// GetAllCart implements usercart.UserCartServiceInterface.
func (*UserCartService) GetAllCart() ([]usercart.UserCartCore, error) {
	panic("unimplemented")
}

// GetCartById implements usercart.UserCartServiceInterface.
func (service *UserCartService) GetCartById(id uint) (usercart.UserCartCore, error) {
	result, err := service.UserCartData.SelectCartById(id)
	return result, err
}

// UpdateCart implements usercart.UserCartServiceInterface.
func (*UserCartService) UpdateCart(id uint, input usercart.UserCartCore) error {
	panic("unimplemented")
}

func New(repo usercart.UserCartDataInterface) usercart.UserCartServiceInterface {
	return &UserCartService{
		UserCartData: repo,
		validate:     &validator.Validate{},
	}
}
