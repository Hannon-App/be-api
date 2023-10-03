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

	var fileName string
	file, header, errFile := c.Request().FormFile("images")

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

	// errUp := helper.Uploader.UploadFile(file, header.Filename)

	// if errUp != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	// }
	var tenantCore = TenantRequestToCore(userInput)
	err := handler.tenantService.Create(tenantCore, file, fileName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "operation success, data added", nil))
}
