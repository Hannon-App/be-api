package data

import (
	"Hannon-app/features/items"
	"errors"

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

func (repo *ItemQuery) ReadAll() ([]items.ItemCore, error) {
	var itemData []Item
	tx := repo.db.Find(&itemData)
	if tx.Error != nil {
		return nil, tx.Error
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
	return itemCore, nil
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

func (repo *ItemQuery) Insert(input items.ItemCore) (items.ItemCore, error) {

	itemGorm := ItemCoreToModel(input)
	tx := repo.db.Create(&itemGorm) // proses query insert
	if tx.Error != nil {
		return items.ItemCore{}, tx.Error
	}
	return ModelToCore(itemGorm), nil
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
