package handler

import (
	"Hannon-app/app/middlewares"
	_item "Hannon-app/features/items/handler"
	"Hannon-app/features/tenants"
	"Hannon-app/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TenantHandler struct {
	tenantService tenants.TenantServiceInterface
}

func New(service tenants.TenantServiceInterface) *TenantHandler {
	return &TenantHandler{
		tenantService: service,
	}
}

func (handler *TenantHandler) Insert(c echo.Context) error {
	var userInput TenantRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBindData.Error(), nil))
	}

	// Handling image upload
	imageFile, imageHeader, errImageFile := c.Request().FormFile("images")
	if errImageFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrReadingFile.Error(), nil))
	}
	imageName := strings.ReplaceAll(imageHeader.Filename, " ", "_")

	// Handling ID card upload
	idCardFile, idCardHeader, errIDCardFile := c.Request().FormFile("id_card")
	if errIDCardFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrReadingFile.Error(), nil))
	}
	idCardName := strings.ReplaceAll(idCardHeader.Filename, " ", "_")

	var tenantCore = TenantRequestToCore(userInput)
	err := handler.tenantService.Create(tenantCore, imageFile, idCardFile, imageName, idCardName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success register tenant", nil))
}

func (handler *TenantHandler) Login(c echo.Context) error {
	userInput := new(TenantLogin)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}

	dataLogin, token, err := handler.tenantService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = TenantLoginResponse{
		ID:    dataLogin.ID,
		Name:  dataLogin.Name,
		Role:  dataLogin.Role,
		Token: token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *TenantHandler) GetAllTenant(c echo.Context) error {
	addressFilter := c.QueryParam("location")

	result, err := handler.tenantService.ReadAll(addressFilter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrReadData.Error(), nil))
	}
	var data []TenantResponse
	for _, value := range result {
		if addressFilter != "" && value.Address != addressFilter {
			continue
		}
		data = append(data, TenantResponse{
			ID:        value.ID,
			Name:      value.Name,
			Address:   value.Address,
			Email:     value.Email,
			Phone:     value.Phone,
			Images:    value.Images,
			OpenTime:  value.OpenTime,
			CloseTime: value.CloseTime,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", data))
}

func (handler *TenantHandler) GetTenantItems(c echo.Context) error {
	id := c.Param("tenant_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}

	result, err := handler.tenantService.ReadAllTenantItems(uint(idParam))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		}
	}
	var tenantItems []TenantItemResponse
	for _, tenant := range result {
		var items []_item.ItemResponseAll
		for _, item := range tenant.Items {
			items = append(items, _item.ItemResponseAll{
				ID:               item.ID,
				Name:             item.Name,
				Stock:            item.Stock,
				Rent_Price:       item.Rent_Price,
				Image:            item.Image,
				Description_Item: item.Description_Item,
				Broke_Cost:       item.Broke_Cost,
				Lost_Cost:        item.Lost_Cost,
				Status:           item.Status,
			})
		}
		tenantItems = append(tenantItems, TenantItemResponse{
			ID:     tenant.ID,
			Name:   tenant.Name,
			Images: tenant.Images,
			Items:  items,
		})
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read tenant items data", tenantItems))

}

func (handler *TenantHandler) DeleteTenant(c echo.Context) error {
	id := c.Param("tenant_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	admin := middlewares.ExtractTokenAdminId(c)
	if admin != 1 {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, helpers.ErrForbiddenAccess.Error(), nil))
	}
	err := handler.tenantService.Remove(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Delete Tenant Successfully", nil))
}

func (handler *TenantHandler) GetTenantById(c echo.Context) error {
	id := c.Param("tenant_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, helpers.ErrBadRequest.Error(), nil))
	}
	result, err := handler.tenantService.ReadTenantById(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, helpers.ErrInternalServer.Error(), nil))
	}

	resultResponse := TenantResponse{
		ID:        result.ID,
		Name:      result.Name,
		Address:   result.Address,
		Email:     result.Email,
		Phone:     result.Phone,
		Images:    result.Images,
		OpenTime:  result.OpenTime,
		CloseTime: result.CloseTime,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read tenant data", resultResponse))

}
