package handler

import "Hannon-app/features/items"

type ItemRequest struct {
	Name             string `json:"name" validate:"required"`
	Stock            uint   `json:"stock" validate:"required,gte=0"`
	Rent_Price       uint   `json:"rent_price" validate:"required,gte=0"`
	Image            string `json:"image"`
	Description_Item string `json:"description_item"`
	Broke_Cost       uint   `json:"broke_cost" validate:"required,gte=0"`
	Lost_Cost        uint   `json:"lost_cost" validate:"required,gte=0"`
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
