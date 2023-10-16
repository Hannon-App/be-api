package handler

import (
	"Hannon-app/features/payments"
	"Hannon-app/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	PaymentService payments.PaymentServiceInterface
}

func New(service payments.PaymentServiceInterface) *PaymentHandler {
	return &PaymentHandler{
		PaymentService: service,
	}
}

func (handler *PaymentHandler) CreateVirtualAccount(c echo.Context) error {
	var paymentData VARequest
	errBind := c.Bind(&paymentData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	var VACore = payments.VirtualAccountObjectCore{
		ExternalID: paymentData.ExternalID,
		BankCode:   paymentData.BankCode,
		Name:       paymentData.Name,
	}
	err := handler.PaymentService.AddVA(VACore)
	if err != nil {
		if err.Error() != "" && strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}
