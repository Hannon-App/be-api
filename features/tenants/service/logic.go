package service

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/tenants"
	"mime/multipart"
)

type TenantService struct {
	tenantData tenants.TenantDataInterface
}

// Create implements tenants.TenantServiceInterface.
func (service *TenantService) Create(input tenants.TenantCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	err := service.tenantData.Register(input, fileImages, fileID, filenameImages, filenameID)
	return err
}

// Edit implements tenants.TenantServiceInterface.
func (*TenantService) Edit(input tenants.TenantCore) error {
	panic("unimplemented")
}

// Login implements tenants.TenantServiceInterface.
func (service *TenantService) Login(email string, password string) (dataLogin tenants.TenantCore, token string, err error) {
	dataLogin, err = service.tenantData.Login(email, password)
	if err != nil {
		return tenants.TenantCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.ID, dataLogin.ID)
	if err != nil {
		return tenants.TenantCore{}, "", err
	}
	return dataLogin, token, nil

}

// ReadAll implements tenants.TenantServiceInterface.
func (service *TenantService) ReadAll(addressFilter string) ([]tenants.TenantCore, error) {
	results, err := service.tenantData.GetAll(addressFilter)
	return results, err
}

// Remove implements tenants.TenantServiceInterface.
func (*TenantService) Remove(id uint) error {
	panic("unimplemented")
}

func New(repo tenants.TenantDataInterface) tenants.TenantServiceInterface {
	return &TenantService{
		tenantData: repo,
	}
}
