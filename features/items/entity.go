package items

import (
	"mime/multipart"
	"time"
)

type ItemCore struct {
	ID               uint
	Name             string
	Stock            uint
	Rent_Price       uint
	Image            string
	Description_Item string
	Broke_Cost       uint
	Lost_Cost        uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type ItemDataInterface interface {
	ReadAll(page, item uint, search_name string) ([]ItemCore, int64, error)
	Delete(tenantID uint, id uint) error
	SelectById(id uint) (ItemCore, error)
	Insert(tenantID uint, input ItemCore, file multipart.File, filename string) error
	UpdateDataItem(tenantID uint, id uint, input ItemCore, file multipart.File, filename string) error
	ReadItemsByTenant(tenantID uint, page, item uint, searchName string) ([]ItemCore, int64, error)
	ReadArchiveItem(tenantID uint) ([]ItemCore, error)
}

type ItemServiceInterface interface {
	GetAllItem(page, item uint, search_name string) ([]ItemCore, bool, error)
	Delete(tenantID uint, id uint) error
	GetById(id uint) (ItemCore, error)
	Create(tenantID uint, input ItemCore, file multipart.File, filename string) error
	Update(tenantID uint, id uint, input ItemCore, file multipart.File, filename string) error
	GetItemsByTenant(tenantID uint, page, item uint, searchName string) ([]ItemCore, bool, error)
	GetArchiveItem(tenantID uint) ([]ItemCore, error)
}
