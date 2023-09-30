package handler

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/users"
	"Hannon-app/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = LoginResponse{
		ID:    dataLogin.ID,
		Token: token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) Add(c echo.Context) error {
	idUser := middlewares.ExtractTokenUserId(c)

	var request UserRequest
	errBind := c.Bind(&request)
	if errBind != nil {
		return helpers.FailedRequest(c, "error bind data"+errBind.Error(), nil)
	}
	link, errLink := helpers.UploadImage(c)
	if errLink != nil {
		return helpers.FailedRequest(c, errLink.Error(), nil)
	}

	fmt.Println(request)
	entity := RequestToCore(request)
	entity.ID = idUser
	entity.UploadKTP = link
	entity.ProfilePhoto = link
	err := handler.userService.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helpers.FailedRequest(c, err.Error(), nil)
		} else {
			return helpers.InternalError(c, err.Error(), nil)
		}
	}
	return helpers.SuccessCreate(c, "success create reimbursment", nil)
}
