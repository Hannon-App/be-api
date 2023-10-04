package handler

import (
	"Hannon-app/features/tenants"
	"Hannon-app/helpers"
	"net/http"
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
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid"+errBind.Error(), nil))
	}

	// Handling image upload
	imageFile, imageHeader, errImageFile := c.Request().FormFile("images")
	if errImageFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading image file: "+errImageFile.Error(), nil))
	}
	imageName := strings.ReplaceAll(imageHeader.Filename, " ", "_")

	// Handling ID card upload
	idCardFile, idCardHeader, errIDCardFile := c.Request().FormFile("id_card")
	if errIDCardFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading ID card file: "+errIDCardFile.Error(), nil))
	}
	idCardName := strings.ReplaceAll(idCardHeader.Filename, " ", "_")

	var tenantCore = TenantRequestToCore(userInput)
	err := handler.tenantService.Create(tenantCore, imageFile, idCardFile, imageName, idCardName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success, data added", nil))
}
