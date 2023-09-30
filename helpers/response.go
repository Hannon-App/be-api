package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(code int, message string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type FindAllMapResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	NextPage bool        `json:"next"`
	Data     interface{} `json:"data,omitempty"`
}

func FindAllWebResponse(code int, message string, data interface{}, nextPage bool) FindAllMapResponse {
	return FindAllMapResponse{
		Code:     code,
		Message:  message,
		Data:     data,
		NextPage: nextPage,
	}
}

func FailedNotFound(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusNotFound, map[string]any{
		"status":  "fail",
		"message": message,
		"data":    data,
	})
}

func Found(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusFound, map[string]any{
		"status":  "succes",
		"message": message,
		"data":    data,
	})
}

func FailedRequest(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"status":  "fail",
		"message": message,
		"data":    data,
	})
}

func SuccessCreate(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func UnAutorization(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusUnauthorized, map[string]any{
		"status":  "fail",
		"message": message,
		"data":    data,
	})
}

func Forbidden(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusForbidden, map[string]any{
		"status":  "fail",
		"message": message,
		"data":    data,
	})
}

func InternalError(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusInternalServerError, map[string]any{
		"status":  "fail",
		"message": message,
		"data":    data,
	})
}
