package handler

import (
	"Hannon-app/app/config"
	"Hannon-app/app/middlewares"
	"Hannon-app/features/rents"
	"Hannon-app/helpers"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	rentService rents.RentServiceInterface
	rentData    rents.RentDataInterface
	config      *config.AppConfig
}

func New(service rents.RentServiceInterface, repo rents.RentDataInterface) *RentHandler {
	return &RentHandler{
		rentService: service,
		rentData:    repo,
	}
}

func (handler *RentHandler) CreateRent(c echo.Context) error {
	var rentData RentRequest
	errBind := c.Bind(&rentData)
	userID := middlewares.ExtractTokenUserId(c)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	layoutFormat := "2006-01-02 15:04:05"
	valueStartDate := rentData.StartDate
	valueEndDate := rentData.EndDate
	startDate, errStart := time.Parse(layoutFormat, valueStartDate)
	endDate, errEnd := time.Parse(layoutFormat, valueEndDate)

	// Error handling for time parsing
	if errStart != nil || errEnd != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid date format", nil))
	}

	var rentCore = rents.RentCore{
		StartDate:     startDate,
		EndDate:       endDate,
		Status:        rentData.Status,
		TotalPrice:    rentData.TotalPrice,
		Discount:      rentData.Discount,
		PaymentLink:   rentData.PaymentLink,
		InvoiceNumber: rentData.InvoiceNumber,
		IDXendit:      rentData.IDXendit,
		UserID:        userID,
	}

	err := handler.rentService.Add(rentCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *RentHandler) ReadRentById(c echo.Context) error {
	id := c.Param("rent_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	result, err := handler.rentService.ReadById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))

		}
	}
	formattedStartDate := result.StartDate.Format("2006-01-02 15:04:05")
	formattedEndDate := result.EndDate.Format("2006-01-02 15:04:05")
	resultResponse := RentResponse{
		ID:            result.ID,
		StartDate:     formattedStartDate,
		EndDate:       formattedEndDate,
		Status:        result.Status,
		TotalPrice:    result.TotalPrice,
		Discount:      result.Discount,
		PaymentLink:   result.PaymentLink,
		InvoiceNumber: result.InvoiceNumber,
		UserID:        result.UserID,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}

func (handler *RentHandler) UpdatebyId(c echo.Context) error {
	var input RentRequest
	id := c.Param("rent_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	rentCore := RentRequestToCore(input)
	err := handler.rentService.UpdatebyId(uint(idConv), rentCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update rent data", nil))
}

func (handler *RentHandler) Payment(c echo.Context) error {
	var data RentRequest
	idParam := c.Param("rent_id")
	id, errConv := strconv.Atoi(idParam)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "wrong id", nil))
	}
	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	userID := middlewares.ExtractTokenUserId(c)
	data.UserID = userID

	err := handler.rentService.AcceptPayment(uint(id), data.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, string(helpers.ErrCheckoutFail.Error()), nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success create invoice", nil))
}

func (handler *RentHandler) Callback(c echo.Context) error {
	req := c.Request()
	headers := req.Header

	callBackToken := headers.Get("X-Callback-Token")

	if callBackToken != handler.config.CallbackKey {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	var callback CallBackRequest
	c.Bind(&callback)

	callBackData := rents.RentCore{
		Status:        callback.Status,
		InvoiceNumber: callback.InvoiceNumber,
	}

	err := handler.rentService.Callback(callBackData)

	return err
}
