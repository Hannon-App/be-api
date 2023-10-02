package handler

import "Hannon-app/features/items"

type ItemRequest struct {
	ID               uint   `gorm:"column:id;primaryKey"`
	Name             string `gorm:"name;not null"`
	Stock            uint   `gorm:"stock;not null"`
	Rent_Price       uint   `gorm:"rent_price;not null"`
	Image            string `gorm:"image;not null"`
	Description_Item string `gorm:"description_item;not null"`
	Broke_Cost       uint   `gorm:"broke_cost;not null"`
	Lost_Cost        uint   `gorm:"lost_cost;not null"`
	Package_Item     string `gorm:"package_item;not null"`
}

func RequestToCore(input ItemRequest) items.ItemCore {
	return items.ItemCore{
		Name:             input.Name,
		Stock:            input.Stock,
		Rent_Price:       input.Rent_Price,
		Image:            input.Image,
		Description_Item: input.Description_Item,
		Broke_Cost:       input.Broke_Cost,
		Lost_Cost:        input.Lost_Cost,
	}
}

type ItemUpdateRequest struct {
	Name             string `json:"name"`
	Stock            uint   `json:"stock"`
	Rent_Price       uint   `json:"rent_price"`
	Image            string `json:"image"`
	Description_Item string `json:"description_item"`
	Broke_Cost       uint   `json:"broke_cost"`
	Lost_Cost        uint   `json:"lost_cost"`
}

func ItemUpdateRequestToCore(input ItemUpdateRequest) items.ItemCore {
	return items.ItemCore{
		Name:             input.Name,
		Stock:            input.Stock,
		Rent_Price:       input.Rent_Price,
		Image:            input.Image,
		Description_Item: input.Description_Item,
		Broke_Cost:       input.Broke_Cost,
		Lost_Cost:        input.Lost_Cost,
	}

}
