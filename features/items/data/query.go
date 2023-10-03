package data

import (
	"Hannon-app/features/items"
	"Hannon-app/helpers"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type ItemQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) items.ItemDataInterface {
	return &ItemQuery{
		db: db,
	}
}

func (repo *ItemQuery) ReadAll(page, itemPerPage uint, searchName string) ([]items.ItemCore, int64, error) {
	var itemData []Item
	var totalCount int64

	if page == 0 && itemPerPage == 0 {
		tx := repo.db

		if searchName != "" {
			tx = tx.Where("name LIKE ?", "%"+searchName+"%")
		}
		tx.Find(&itemData)
	} else {

		offset := int((page - 1) * itemPerPage)

		query := repo.db.Offset(offset).Limit(int(itemPerPage))

		if searchName != "" {
			query = query.Where("name LIKE ?", "%"+searchName+"%")
		}

		tx := query.Find(&itemData)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
	}

	var itemCore []items.ItemCore
	for _, value := range itemData {
		itemCore = append(itemCore, items.ItemCore{
			ID:               value.ID,
			Name:             value.Name,
			Stock:            value.Stock,
			Rent_Price:       value.Rent_Price,
			Image:            value.Image,
			Description_Item: value.Description_Item,
			Broke_Cost:       value.Broke_Cost,
			Lost_Cost:        value.Lost_Cost,
		})
	}

	repo.db.Model(&Item{}).Count(&totalCount)

	return itemCore, totalCount, nil
}

func (repo *ItemQuery) Delete(id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&Item{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *ItemQuery) SelectById(id uint) (items.ItemCore, error) {
	var result Item
	tx := repo.db.Find(&result, id)
	if tx.Error != nil {
		return items.ItemCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return items.ItemCore{}, errors.New("no row affected")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

func (repo *ItemQuery) Insert(input items.ItemCore, file multipart.File, filename string) error {

	var itemModel = ItemCoreToModel(input)

	if filename == "default.png" {
		itemModel.Image = filename
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		itemModel.Image = nameGen + filename
		errUp := helpers.Uploader.UploadFile(file, itemModel.Image)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Create(&itemModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ItemQuery) UpdateDataItem(id uint, input items.ItemCore) (items.ItemCore, error) {
	itemGorm := ItemCoreToModel(input)
	tx := repo.db.Model(&Item{}).Where("id = ?", id).Updates(itemGorm)
	if tx.Error != nil {
		return items.ItemCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return items.ItemCore{}, errors.New("item not found")
	}
	return ModelToCore(itemGorm), nil
}
