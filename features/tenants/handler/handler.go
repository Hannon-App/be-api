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
