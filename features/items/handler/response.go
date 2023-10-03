package handler

import "Hannon-app/features/items"

type ItemResponse struct {
	Name             string `json:"name"`
	Stock            uint   `json:"stock"`
	Rent_Price       uint   `json:"rent_price"`
	Image            string `json:"image"`
	Description_Item string `json:"description_item"`
	Broke_Cost       uint   `json:"broke_cost"`
	Lost_Cost        uint   `json:"lost_cost"`
}

type ItemResponseAll struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Stock            uint   `json:"stock"`
	Rent_Price       uint   `json:"rent_price"`
	Image            string `json:"image"`
	Description_Item string `json:"description_item"`
	Broke_Cost       uint   `json:"broke_cost"`
	Lost_Cost        uint   `json:"lost_cost"`
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

type ItemCreateResponse struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Stock            uint   `json:"stock"`
	Rent_Price       uint   `json:"rent_price"`
	Image            string `json:"image"`
	Description_Item string `json:"description_item"`
	Broke_Cost       uint   `json:"broke_cost"`
	Lost_Cost        uint   `json:"lost_cost"`
}
