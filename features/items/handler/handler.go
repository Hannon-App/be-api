package handler

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/items"
	"Hannon-app/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
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

	var pageConv, itemConv int
	var errPageConv, errItemConv error

	page := c.QueryParam("page")
	if page != "" {
		pageConv, errPageConv = strconv.Atoi(page)
		if errPageConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}
	item := c.QueryParam("itemPerPage")
	if item != "" {
		itemConv, errItemConv = strconv.Atoi(item)
		if errItemConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}

	search_name := c.QueryParam("searchName")

	result, next, err := handler.itemService.GetAllItem(uint(pageConv), uint(itemConv), search_name)
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
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "success read data", itemResponse, next))
}

func (handler *ItemHandler) DeleteItem(c echo.Context) error {

	tenantID, er := middlewares.ExtractTokenTenant(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	id := c.Param("item_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	err := handler.itemService.Delete(tenantID, uint(idConv))
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

	tenantID, er := middlewares.ExtractTokenTenant(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var itemInput ItemRequest
	errBind := c.Bind(&itemInput)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	var fileName string
	file, header, errFile := c.Request().FormFile("image")

	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			fileName = "default.png"
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid "+errFile.Error(), nil))
		}
	}

	if fileName == "" {
		fileName = strings.ReplaceAll(header.Filename, " ", "_")
	}

	validate := validator.New()
	if err := validate.Struct(itemInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	itemCore := RequestToCore(itemInput)
	err := handler.itemService.Create(tenantID, itemCore, file, fileName)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *ItemHandler) UpdateItemByID(c echo.Context) error {

	tenantID, er := middlewares.ExtractTokenTenant(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var itemInput ItemUpdateRequest

	id := c.Param("item_id")

	itemID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid item ID", nil))
	}

	if err := c.Bind(&itemInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	var fileName string
	file, header, errFile := c.Request().FormFile("image")

	if errFile != nil {
		if strings.Contains(errFile.Error(), "no such file") {
			fileName = "default.png"
		} else {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid "+errFile.Error(), nil))
		}
	}

	if fileName == "" {
		fileName = strings.ReplaceAll(header.Filename, " ", "_")
	}

	validate := validator.New()
	if err := validate.Struct(itemInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	itemCore := ItemUpdateRequestToCore(itemInput)
	if err := handler.itemService.Update(tenantID, uint(itemID), itemCore, file, fileName); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "item not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating item: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success update data", nil))
}

func (handler *ItemHandler) GetAllItemsTenant(c echo.Context) error {

	tenantID, err := middlewares.ExtractTokenTenant(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, err.Error(), nil))
	}

	var pageConv, itemConv int
	var errPageConv, errItemConv error

	page := c.QueryParam("page")
	if page != "" {
		pageConv, errPageConv = strconv.Atoi(page)
		if errPageConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}
	item := c.QueryParam("itemPerPage")
	if item != "" {
		itemConv, errItemConv = strconv.Atoi(item)
		if errItemConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}

	search_name := c.QueryParam("searchName")

	result, next, err := handler.itemService.GetItemsByTenant(tenantID, uint(pageConv), uint(itemConv), search_name)
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
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "success read data", itemResponse, next))
}
