package handler

import (
	"Hannon-app/features/users"
	"Hannon-app/helpers"
	"strconv"

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

func (handler *UserHandler) AddUser(c echo.Context) error {
	link, errUpload := helpers.UploadImage(c)
	if errUpload != nil {
		return helpers.FailedRequest(c, errUpload.Error(), nil)
	}

	var input UserRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return helpers.FailedNotFound(c, "error binding", nil)
	}
	entity := RequestToCore(input)
	entity.ProfilePhoto = link
	err := handler.userService.Add(entity)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helpers.FailedRequest(c, err.Error(), nil)
		} else {
			return helpers.InternalError(c, err.Error(), nil)
		}
	}
	return helpers.SuccessWithOutData(c, "success create Users")
}

func (handler *UserHandler) GetUsertById(c echo.Context) error {
	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	result, err := handler.userService.GetUserById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	resultResponse := UserCoreToResponse(result)

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error binding data", nil))
	}

	// Mapping data dari UserRequest ke struct Core
	var userCore uint
	err := handler.userService.Deletebyid(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "User not found", nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error Deleted data user", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "User Deleted successfully", nil))
}
