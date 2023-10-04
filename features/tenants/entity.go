package tenants

import (
	"mime/multipart"
	"time"
)

type TenantCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Images    string
	Address   string
	IDcard    string
	OpenTime  string
	CloseTime string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type TenantDataInterface interface {
	Login(email string, password string) (dataLogin TenantCore, err error)
	Register(input TenantCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	GetAll() ([]TenantCore, error)
	Update(input TenantCore) error
	Delete(id uint) error
}

type TenantServiceInterface interface {
	Login(email string, password string) (dataLogin TenantCore, token string, err error)
	Create(input TenantCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error
	ReadAll() ([]TenantCore, error)
	Edit(input TenantCore) error
	Remove(id uint) error
}
