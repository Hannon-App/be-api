package data

import (
	"Hannon-app/features/tenants"
	"Hannon-app/helpers"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type TenantQuery struct {
	db *gorm.DB
	// dataLogin tenants.TenantCore
}

// Delete implements tenants.TenantDataInterface.
func (*TenantQuery) Delete(id uint) error {
	panic("unimplemented")
}

// GetAll implements tenants.TenantDataInterface.
func (*TenantQuery) GetAll() ([]tenants.TenantCore, error) {
	panic("unimplemented")
}

// Login implements tenants.TenantDataInterface.
func (*TenantQuery) Login(email string, password string) (dataLogin tenants.TenantCore, err error) {
	panic("unimplemented")
}

// Register implements tenants.TenantDataInterface.
func (repo *TenantQuery) Register(input tenants.TenantCore, fileImages multipart.File, fileID multipart.File, filenameImages string, filenameID string) error {
	var tenantModel = TenantCoreToModel(input)

	hash, errHass := helpers.HashPassword(tenantModel.Password)
	if errHass != nil {
		return errHass
	}
	tenantModel.Password = hash

	if filenameImages == "default.png" {
		tenantModel.Images = filenameImages
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		tenantModel.Images = nameGen + filenameImages
		errUp := helpers.Uploader.UploadFile(fileImages, tenantModel.Images)

		if errUp != nil {
			return errUp
		}
	}

	if filenameID == "default.png" {
		tenantModel.IDcard = filenameID
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		tenantModel.IDcard = nameGen + filenameID
		errUp := helpers.Uploader.UploadFile(fileID, tenantModel.IDcard)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Create(&tenantModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

// Update implements tenants.TenantDataInterface.
func (*TenantQuery) Update(input tenants.TenantCore) error {
	panic("unimplemented")
}

func New(db *gorm.DB) tenants.TenantDataInterface {
	return &TenantQuery{
		db: db,
	}
}
