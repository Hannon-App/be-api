package usercart

type UserCartCore struct {
	ID        uint
	UserID    uint
	TenantID  uint
	CartItems []CartItemCore
}

type CartItemCore struct {
	ID         uint
	UserCartID uint
	ItemID     uint
	Quantity   uint
}

type UserCartDataInterface interface {
	CreateCart(input UserCartCore) error
	UpdateCart(id uint, input UserCartCore) error
	DeleteCart(id uint) error
	SelectAllCart() ([]UserCartCore, error)
	SelectCartById(id uint) (UserCartCore, error)
}

type UserCartServiceInterface interface {
	CreateCart(input UserCartCore) error
	UpdateCart(id uint, input UserCartCore) error
	DeleteCart(id uint) error
	GetAllCart() ([]UserCartCore, error)
	GetCartById(id uint) (UserCartCore, error)
}
