package handler

import (
	"Hannon-app/features/items"
	"Hannon-app/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService items.ItemServiceInterface
}

func New(service items.ItemServiceInterface) *ItemHandler {
	return &ItemHandler{
		itemService: service,
	}
}

func (handler *ItemHandler) GetAll(c echo.Context) error {

	result, err := handler.itemService.GetAllItem()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var itemResponse []ItemResponseAll
	for _, value := range result {
		itemResponse = append(itemResponse, ItemResponseAll{
			ID:               value.ID,
			Name:             value.Name,
			Stock:            value.Stock,
			Rent_Price:       value.Rent_Price,
			Image:            value.Image,
			Description_Item: value.Description_Item,
			Broke_Cost:       value.Broke_Cost,
			Lost_Cost:        value.Lost_Cost,
		})

	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", itemResponse))
}

func (handler *ItemHandler) DeleteItem(c echo.Context) error {

	id := c.Param("item_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	err := handler.itemService.Delete(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "operation failed, requested resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *ItemHandler) GetItemByID(c echo.Context) error {
	id := c.Param("item_id")

	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	result, err := handler.itemService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "operation failed, requested resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}

	resultResponse := ItemCoreToResponseAll(result)

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get item data", resultResponse))
}
