package router

import (
	_adminData "Hannon-app/features/admins/data"
	_adminHandler "Hannon-app/features/admins/handler"
	_adminService "Hannon-app/features/admins/service"
	_userData "Hannon-app/features/users/data"
	_userHandler "Hannon-app/features/users/handler"
	_userService "Hannon-app/features/users/service"

	_itemData "Hannon-app/features/items/data"
	_itemHandler "Hannon-app/features/items/handler"
	_itemService "Hannon-app/features/items/service"
	"Hannon-app/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	ItemData := _itemData.New(db)
	ItemService := _itemService.New(ItemData)
	ItemHandlerAPI := _itemHandler.New(ItemService)

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

	c.GET("/items", ItemHandlerAPI.GetAll)
	c.DELETE("/items/:item_id", ItemHandlerAPI.DeleteItem)
	c.GET("/items/:item_id", ItemHandlerAPI.GetItemByID)
	c.POST("/items", ItemHandlerAPI.CreateItem)
	c.PUT("/items/:item_id", ItemHandlerAPI.UpdateItemByID)

	//Admin
	c.POST("/admin", AdminHandlerAPI.Login)

}
