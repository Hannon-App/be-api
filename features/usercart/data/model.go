package data

import (
	"Hannon-app/features/usercart"

	"gorm.io/gorm"
)

type UserCart struct {
	gorm.Model
	UserID    uint
	TenantID  uint
	CartItems []CartItem
}

type CartItem struct {
	gorm.Model
	UserCartID uint
	ItemID     uint
	Quantity   uint
}

func ToUserCartCore(userCart UserCart) usercart.UserCartCore {
	cartItemsCore := make([]usercart.CartItemCore, len(userCart.CartItems))
	for i, item := range userCart.CartItems {
		cartItemsCore[i] = usercart.CartItemCore{
			ID:         item.ID,
			UserCartID: item.UserCartID,
			ItemID:     item.ItemID,
			Quantity:   item.Quantity,
		}
	}
	return usercart.UserCartCore{
		ID:        userCart.ID,
		UserID:    userCart.UserID,
		TenantID:  userCart.TenantID,
		CartItems: cartItemsCore,
	}
}

func ToCartItemCore(cartItem CartItem) usercart.CartItemCore {
	return usercart.CartItemCore{
		ID:         cartItem.ID,
		UserCartID: cartItem.UserCartID,
		ItemID:     cartItem.ItemID,
		Quantity:   cartItem.Quantity,
	}
}

func FromUserCartCore(userCartCore usercart.UserCartCore) UserCart {
	cartItems := make([]CartItem, len(userCartCore.CartItems))
	for i, itemCore := range userCartCore.CartItems {
		cartItems[i] = CartItem{
			Model:      gorm.Model{ID: itemCore.ID},
			UserCartID: itemCore.UserCartID,
			ItemID:     itemCore.ItemID,
			Quantity:   itemCore.Quantity,
		}
	}
	return UserCart{
		Model:     gorm.Model{ID: userCartCore.ID},
		UserID:    userCartCore.UserID,
		TenantID:  userCartCore.TenantID,
		CartItems: cartItems,
	}
}

func FromCartItemCore(cartItemCore usercart.CartItemCore) CartItem {
	return CartItem{
		Model:      gorm.Model{ID: cartItemCore.ID},
		UserCartID: cartItemCore.UserCartID,
		ItemID:     cartItemCore.ItemID,
		Quantity:   cartItemCore.Quantity,
	}
}
