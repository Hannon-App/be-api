package router

import (
	_userData "Hannon-app/features/users/data"
	_userHandler "Hannon-app/features/users/handler"
	_userService "Hannon-app/features/users/service"
	"Hannon-app/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Welcome to HannonApp!", nil))
	})

	c.POST("/login", UserHandlerAPI.Login)
}
