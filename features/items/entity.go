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
	ReadAll() ([]ItemCore, error)
	Delete(id uint) error
	SelectById(id uint) (ItemCore, error)
}

type ItemServiceInterface interface {
	GetAllItem() ([]ItemCore, error)
	Delete(id uint) error
	GetById(id uint) (ItemCore, error)
}
