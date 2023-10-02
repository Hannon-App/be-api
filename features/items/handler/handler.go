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

func (handler *ItemHandler) CreateItem(c echo.Context) error {
	itemInput := new(ItemRequest)
	errBind := c.Bind(&itemInput)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	itemCore := RequestToCore(*itemInput)
	result, err := handler.itemService.Create(itemCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	Response := ItemCreateResponse{
		ID:               result.ID,
		Name:             result.Name,
		Stock:            result.Stock,
		Rent_Price:       result.Rent_Price,
		Image:            result.Image,
		Description_Item: result.Description_Item,
		Broke_Cost:       result.Broke_Cost,
		Lost_Cost:        result.Lost_Cost,
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", Response))
}

func (handler *ItemHandler) UpdateItemByID(c echo.Context) error {
	idItemStr := c.Param("item_id")
	idItem, errItem := strconv.Atoi(idItemStr)
	if errItem != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "item id invalid", nil))
	}

	itemInput := new(ItemUpdateRequest)
	errBind := c.Bind(itemInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	itemCore := ItemUpdateRequestToCore(*itemInput)
	result, err := handler.itemService.Update(uint(idItem), itemCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
	}

	updateResponse := ItemResponse{
		Name:             result.Name,
		Stock:            result.Stock,
		Rent_Price:       result.Rent_Price,
		Image:            result.Image,
		Description_Item: result.Description_Item,
		Broke_Cost:       result.Broke_Cost,
		Lost_Cost:        result.Lost_Cost,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update data", updateResponse))
}
