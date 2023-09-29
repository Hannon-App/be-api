package router

import (
	_adminData "Hannon-app/features/admins/data"
	_adminHandler "Hannon-app/features/admins/handler"
	_adminService "Hannon-app/features/admins/service"
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

	AdminData := _adminData.New(db)
	AdminService := _adminService.New(AdminData)
	AdminHandlerAPI := _adminHandler.New(AdminService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Welcome to HannonApp!", nil))
	})

	//Users
	c.POST("/login", UserHandlerAPI.Login)

	//Admin
	c.POST("/admin", AdminHandlerAPI.Login)
}
