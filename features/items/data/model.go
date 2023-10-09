package data

import (
	"Hannon-app/features/items"
	_cartData "Hannon-app/features/usercart/data"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name             string `gorm:"name;not null"`
	Stock            uint   `gorm:"stock;not null"`
	Rent_Price       uint   `gorm:"rent_price;not null"`
	Image            string `gorm:"image;not null"`
	Description_Item string `gorm:"description_item;not null"`
	Broke_Cost       uint   `gorm:"broke_cost;not null"`
	Lost_Cost        uint   `gorm:"lost_cost;not null"`
	TenantID         uint
	Status           string `gorm:"default:available"`
	UserCart         []_cartData.CartItem
}

func ItemCoreToModel(input items.ItemCore) Item {
	var itemModel = Item{
		Model:            gorm.Model{},
		Name:             input.Name,
		Stock:            input.Stock,
		Rent_Price:       input.Rent_Price,
		Image:            input.Image,
		Description_Item: input.Description_Item,
		Broke_Cost:       input.Broke_Cost,
		Lost_Cost:        input.Lost_Cost,
		Status:           input.Status,
	}
	return itemModel
}

func ModelToCore(input Item) items.ItemCore {
	var itemCore = items.ItemCore{
		ID:               input.ID,
		Name:             input.Name,
		Stock:            input.Stock,
		Rent_Price:       input.Rent_Price,
		Image:            input.Image,
		Description_Item: input.Description_Item,
		Broke_Cost:       input.Broke_Cost,
		Lost_Cost:        input.Lost_Cost,
		Status:           input.Status,
	}
	return itemCore
}
