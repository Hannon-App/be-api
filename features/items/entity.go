package items

import "mime/multipart"

type ItemCore struct {
	ID               uint
	Name             string
	Stock            uint
	Rent_Price       uint
	Image            string
	Description_Item string
	Broke_Cost       uint
	Lost_Cost        uint
}

type ItemDataInterface interface {
	ReadAll(page, item uint, search_name string) ([]ItemCore, int64, error)
	Delete(tenantID uint, id uint) error
	SelectById(id uint) (ItemCore, error)
	Insert(tenantID uint, input ItemCore, file multipart.File, filename string) error
	UpdateDataItem(tenantID uint, id uint, input ItemCore, file multipart.File, filename string) error
}

type ItemServiceInterface interface {
	GetAllItem(page, item uint, search_name string) ([]ItemCore, bool, error)
	Delete(tenantID uint, id uint) error
	GetById(id uint) (ItemCore, error)
	Create(tenantID uint, input ItemCore, file multipart.File, filename string) error
	Update(tenantID uint, id uint, input ItemCore, file multipart.File, filename string) error
}
