package handler

type ItemRequest struct {
	ID               uint   `gorm:"column:id;primaryKey"`
	Name             string `gorm:"name;not null"`
	Stock            uint   `gorm:"stock;not null"`
	Rent_Price       uint   `gorm:"rent_price;not null"`
	Image            string `gorm:"image;not null"`
	Description_Item string `gorm:"description;not null"`
	Broke_Cost       uint   `gorm:"broke_cost;not null"`
	Lost_Cost        uint   `gorm:"lost_cost;not null"`
	Package_Item     string `gorm:"package_item;not null"`
}
