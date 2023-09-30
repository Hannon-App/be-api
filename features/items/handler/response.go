package handler

import "Hannon-app/features/items"

type ItemResponseAll struct {
	ID               uint   `gorm:"column:id;primaryKey"`
	Name             string `gorm:"name;not null"`
	Stock            uint   `gorm:"stock;not null"`
	Rent_Price       uint   `gorm:"rent_price;not null"`
	Image            string `gorm:"image;not null"`
	Description_Item string `gorm:"description;not null"`
	Broke_Cost       uint   `gorm:"broke_cost;not null"`
	Lost_Cost        uint   `gorm:"lost_cost;not null"`
}

func ItemCoreToResponseAll(input items.ItemCore) ItemResponseAll {
	var itemResp = ItemResponseAll{
		ID:               input.ID,
		Name:             input.Name,
		Stock:            input.Stock,
		Rent_Price:       input.Rent_Price,
		Image:            input.Image,
		Description_Item: input.Description_Item,
		Broke_Cost:       input.Broke_Cost,
		Lost_Cost:        input.Lost_Cost,
	}
	return itemResp
}
