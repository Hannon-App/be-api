package service

import (
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
func (*TenantService) Login(email string, password string) (dataLogin tenants.TenantCore, token string, err error) {
	panic("unimplemented")
}

// ReadAll implements tenants.TenantServiceInterface.
func (*TenantService) ReadAll() ([]tenants.TenantCore, error) {
	panic("unimplemented")
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
