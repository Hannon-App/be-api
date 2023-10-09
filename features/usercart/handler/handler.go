package handler

import (
	"Hannon-app/features/usercart"
	"Hannon-app/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserCartHandler struct {
	userCartService usercart.UserCartServiceInterface
}

func New(service usercart.UserCartServiceInterface) *UserCartHandler {
	return &UserCartHandler{
		userCartService: service,
	}
}

func (handler *UserCartHandler) CreateUserCart(c echo.Context) error {
	var userInput UserCartRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBindData.Error(), nil))
	}
	var userCartCore = ToUserCartCore(userInput)
	err := handler.userCartService.CreateCart(userCartCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "Item added to cart successfully", nil))
}

func (handler *UserCartHandler) GetCartById(c echo.Context) error {
	id := c.Param("cart_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	result, err := handler.userCartService.GetCartById(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
	}
	convertedItems := make([]CartItemCore, len(result.CartItems))
	for i, item := range result.CartItems {
		convertedItems[i] = CartItemCore{
			ItemID:   item.ItemID,
			Quantity: item.Quantity,
		}
	}

	resultResponse := UserCartResponse{
		ID:       result.ID,
		TenantID: result.TenantID,
		Items:    convertedItems,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read cart data", resultResponse))
}
