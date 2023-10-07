package handler

import "Hannon-app/features/items"

type ItemStatusRequest struct {
	Status string `json:"status"`
}

func ItemStatusToCore(input ItemStatusRequest) items.ItemCore {
	return items.ItemCore{
		Status: input.Status,
	}
}

type ItemRequest struct {
	Name             string `json:"name" form:"name" validate:"required"`
	Stock            uint   `json:"stock" form:"stock" validate:"required"`
	Rent_Price       uint   `json:"rent_price" form:"rent_price" validate:"required"`
	Image            string `json:"image" form:"image"`
	Description_Item string `json:"description_item" form:"description_item"`
	Broke_Cost       uint   `json:"broke_cost" form:"broke_cost" validate:"required"`
	Lost_Cost        uint   `json:"lost_cost" form:"lost_cost" validate:"required"`
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
	Name             string `json:"name" form:"name"`
	Stock            uint   `json:"stock" form:"stock"`
	Rent_Price       uint   `json:"rent_price" form:"rent_price"`
	Image            string `json:"image" form:"image"`
	Description_Item string `json:"description_item" form:"description_item"`
	Broke_Cost       uint   `json:"broke_cost" form:"broke_cost"`
	Lost_Cost        uint   `json:"lost_cost" form:"lost_cost"`
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
