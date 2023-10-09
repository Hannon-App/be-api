package handler

type UserCartResponse struct {
	ID       uint           `json:"id"`
	TenantID uint           `json:"tenant_id"`
	Items    []CartItemCore `json:"items"`
}

type CartItemCore struct {
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}
