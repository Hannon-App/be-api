package data

import (
	_item "Hannon-app/features/items"
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

// GetTenantById implements tenants.TenantDataInterface.
func (repo *TenantQuery) GetTenantById(id uint) (tenants.TenantCore, error) {
	var data Tenant
	tx := repo.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return tenants.TenantCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return tenants.TenantCore{}, errors.New("data not found")
	}

	resultCore := TenantModelToCore(data)
	return resultCore, nil
}

// GetAllTenantItems implements tenants.TenantDataInterface.
func (repo *TenantQuery) GetAllTenantItems(id uint) ([]tenants.TenantCore, error) {
	var tenantData []Tenant

	tx := repo.db.Where("id = ?", id).Preload("Items").Find(&tenantData)
	if tx.Error != nil {
		return []tenants.TenantCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return []tenants.TenantCore{}, errors.New("data not found")
	}

	var tenantsCore []tenants.TenantCore
	for _, value := range tenantData {
		var items []_item.ItemCore
		for _, item := range value.Items {
			items = append(items, _item.ItemCore{
				ID:               item.ID,
				Name:             item.Name,
				Stock:            item.Stock,
				Rent_Price:       item.Rent_Price,
				Image:            item.Image,
				Description_Item: item.Description_Item,
				Broke_Cost:       item.Broke_Cost,
				Lost_Cost:        item.Lost_Cost,
			})
		}

		var tenantCore = tenants.TenantCore{
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
			Items:     items,
		}
		tenantsCore = append(tenantsCore, tenantCore)
	}

	return tenantsCore, nil
}

// Delete implements tenants.TenantDataInterface.
func (repo *TenantQuery) Delete(id uint) error {
	var data Tenant
	tx := repo.db.Where("id = ?", id).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
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
