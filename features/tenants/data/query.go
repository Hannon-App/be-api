package data

import (
	"Hannon-app/features/tenants"
	"Hannon-app/helpers"
	"errors"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type TenantQuery struct {
	db        *gorm.DB
	dataLogin tenants.TenantCore
}

// Delete implements tenants.TenantDataInterface.
func (*TenantQuery) Delete(id uint) error {
	panic("unimplemented")
}

// GetAll implements tenants.TenantDataInterface.
func (repo *TenantQuery) GetAll(addressFilter string) ([]tenants.TenantCore, error) {
	var data []Tenant
	tx := repo.db.Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var TenantCore []tenants.TenantCore
	for _, value := range data {
		if addressFilter != "" && value.Address != addressFilter {
			continue // Skip this entry if the address doesn't match the filter
		}
		TenantCore = append(TenantCore, tenants.TenantCore{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Password:  value.Password,
			Phone:     value.Phone,
			Images:    value.Images,
			Address:   value.Address,
			Role:      value.Role,
			IDcard:    value.IDcard,
			OpenTime:  value.OpenTime,
			CloseTime: value.CloseTime,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		})
	}
	return TenantCore, nil
}

// Login implements tenants.TenantDataInterface.
func (repo *TenantQuery) Login(email string, password string) (dataLogin tenants.TenantCore, err error) {
	var data Tenant
	tx := repo.db.Where("email = ?", email).Find(&data)
	if tx.Error != nil {
		return tenants.TenantCore{}, tx.Error
	}
	check := helpers.CheckPassword(password, data.Password)
	if !check {
		return tenants.TenantCore{}, errors.New("password incorect")
	}
	if tx.RowsAffected == 0 {
		return tenants.TenantCore{}, errors.New("no row affected")
	}
	dataLogin = TenantModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
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
