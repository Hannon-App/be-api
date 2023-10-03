package items

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
	Delete(id uint) error
	SelectById(id uint) (ItemCore, error)
	Insert(input ItemCore) (ItemCore, error)
	UpdateDataItem(id uint, input ItemCore) (ItemCore, error)
}

type ItemServiceInterface interface {
	GetAllItem(page, item uint, search_name string) ([]ItemCore, bool, error)
	Delete(id uint) error
	GetById(id uint) (ItemCore, error)
	Create(input ItemCore) (ItemCore, error)
	Update(id uint, input ItemCore) (ItemCore, error)
}
