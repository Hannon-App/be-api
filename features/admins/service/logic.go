package service

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/admins"
)

type AdminService struct {
	adminData admins.AdminDataInterface
}

// Login implements admins.AdminServiceInterface.
func (service *AdminService) Login(email string, password string) (dataLogin admins.AdminCore, token string, err error) {
	dataLogin, err = service.adminData.Login(email, password)
	if err != nil {
		return admins.AdminCore{}, "", err
	}
	token, err = middlewares.CreateTokenAdmin(dataLogin.ID)
	if err != nil {
		return admins.AdminCore{}, "", err
	}
	return dataLogin, token, nil
}

func New(repo admins.AdminDataInterface) admins.AdminServiceInterface {
	return &AdminService{
		adminData: repo,
	}
}
