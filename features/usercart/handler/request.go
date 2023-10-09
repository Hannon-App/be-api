package handler

import "Hannon-app/features/usercart"

type UserCartRequest struct {
	TenantID uint              `json:"tenant_id"`
	Items    []CartItemRequest `json:"items"`
}

type CartItemRequest struct {
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

func ToUserCartCore(req UserCartRequest) usercart.UserCartCore {
	itemsCore := make([]usercart.CartItemCore, len(req.Items))
	for i, item := range req.Items {
		itemsCore[i] = usercart.CartItemCore{
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
	}
	return usercart.UserCartCore{
		TenantID:  req.TenantID,
		CartItems: []usercart.CartItemCore{},
	}
}
