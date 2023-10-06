package handler

import (
	"Hannon-app/app/middlewares"
	"Hannon-app/features/users"
	"Hannon-app/helpers"
	"strconv"

	"net/http"
	"strings"

	"github.com/go-playground/validator"
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
		Role:  dataLogin.Role,
		Name:  dataLogin.Name,
		Token: token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) AddUser(c echo.Context) error {

	var userInput UserRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid"+errBind.Error(), nil))
	}
	// Handling image upload
	imageFile, imageHeader, errImageFile := c.Request().FormFile("profil_photo")
	if errImageFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading image file: "+errImageFile.Error(), nil))
	}
	imageName := strings.ReplaceAll(imageHeader.Filename, " ", "_")

	// Handling ID card upload
	idCardFile, idCardHeader, errIDCardFile := c.Request().FormFile("ktp_photo")
	if errIDCardFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading ID card file: "+errIDCardFile.Error(), nil))
	}
	idCardName := strings.ReplaceAll(idCardHeader.Filename, " ", "_")

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	var userCore = RequestToCore(userInput)
	err := handler.userService.Add(userCore, imageFile, idCardFile, imageName, idCardName)
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
	adminID, er := middlewares.ExtractTokenAdmin(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	err := handler.userService.Deletebyid(adminID, uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "operation failed, requested resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	uID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var userInput UserRequest
	id := c.Param("user_id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid user ID", nil))
	}

	// Parse and validate the updated user data from the request body.
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	// Handling image upload
	imageFile, imageHeader, errImageFile := c.Request().FormFile("profil_photo")
	if errImageFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading image file: "+errImageFile.Error(), nil))
	}
	imageName := strings.ReplaceAll(imageHeader.Filename, " ", "_")

	// Handling ID card upload
	idCardFile, idCardHeader, errIDCardFile := c.Request().FormFile("ktp_photo")
	if errIDCardFile != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Error reading ID card file: "+errIDCardFile.Error(), nil))
	}
	idCardName := strings.ReplaceAll(idCardHeader.Filename, " ", "_")

	// Create a CoreUser instance with the updated data.
	updatedUser := RequestToCore(userInput)

	// Call the UpdateUser method in the service layer to update the user.
	if err := handler.userService.Update(uID, uint(userID), updatedUser, imageFile, idCardFile, imageName, idCardName); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating user: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success updating user", nil))
}

func (handler *UserHandler) GetAllUser(c echo.Context) error {
	adminID, er := middlewares.ExtractTokenAdmin(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var pageConv, userConv int
	var errPageConv, errUserConv error

	page := c.QueryParam("page")
	if page != "" {
		pageConv, errPageConv = strconv.Atoi(page)
		if errPageConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}
	user := c.QueryParam("userPerPage")
	if user != "" {
		userConv, errUserConv = strconv.Atoi(user)
		if errUserConv != nil {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
		}
	}

	search_name := c.QueryParam("searchName")

	result, next, err := handler.userService.GetAll(adminID, uint(pageConv), uint(userConv), search_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var userResponse []UserResponse
	for _, value := range result {
		userResponse = append(userResponse, UserCoreToResponse(value))

	}
	return c.JSON(http.StatusOK, helpers.FindAllWebResponse(http.StatusOK, "success read data", userResponse, next))
}
