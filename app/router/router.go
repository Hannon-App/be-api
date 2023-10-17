package router

import (
	"Hannon-app/app/config"
	"Hannon-app/app/middlewares"
	_adminData "Hannon-app/features/admins/data"
	_adminHandler "Hannon-app/features/admins/handler"
	_adminService "Hannon-app/features/admins/service"

	_userData "Hannon-app/features/users/data"
	_userHandler "Hannon-app/features/users/handler"
	_userService "Hannon-app/features/users/service"

	_tenantData "Hannon-app/features/tenants/data"
	_tenantHandler "Hannon-app/features/tenants/handler"
	_tenantService "Hannon-app/features/tenants/service"

	_itemData "Hannon-app/features/items/data"
	_itemHandler "Hannon-app/features/items/handler"
	_itemService "Hannon-app/features/items/service"

	_rentData "Hannon-app/features/rents/data"
	_rentHandler "Hannon-app/features/rents/handler"
	_rentService "Hannon-app/features/rents/service"

	_cartData "Hannon-app/features/usercart/data"
	_cartHandler "Hannon-app/features/usercart/handler"
	_cartService "Hannon-app/features/usercart/service"

	_paymentData "Hannon-app/features/payments/data"
	_paymentHandler "Hannon-app/features/payments/handler"
	_paymentService "Hannon-app/features/payments/service"

	"Hannon-app/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo, cfg *config.AppConfig) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	ItemData := _itemData.New(db)
	ItemService := _itemService.New(ItemData)
	ItemHandlerAPI := _itemHandler.New(ItemService)

	AdminData := _adminData.New(db)
	AdminService := _adminService.New(AdminData)
	AdminHandlerAPI := _adminHandler.New(AdminService)

	TenantData := _tenantData.New(db)
	TenantService := _tenantService.New(TenantData)
	TenantHandlerAPI := _tenantHandler.New(TenantService)

	CartData := _cartData.New(db)
	CartService := _cartService.New(CartData)
	CartHandlerAPI := _cartHandler.New(CartService)

	RentData := _rentData.New(db)
	RentService := _rentService.New(RentData, UserData, cfg)
	RentHandlerAPI := _rentHandler.New(RentService, RentData, cfg)

	PaymentData := _paymentData.New(db)
	PaymentService := _paymentService.New(PaymentData, cfg)
	PaymentHandlerAPI := _paymentHandler.New(PaymentService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Welcome to HannonApp!", nil))
	})

	//Users
	c.POST("/login", UserHandlerAPI.Login)
	c.POST("/users", UserHandlerAPI.AddUser)
	c.GET("/users/:user_id", UserHandlerAPI.GetUsertById, middlewares.JWTMiddleware())
	c.DELETE("/users/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	c.PUT("/users/:user_id", UserHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
	c.GET("/users", UserHandlerAPI.GetAllUser, middlewares.JWTMiddleware())

	//items
	c.GET("/items", ItemHandlerAPI.GetAll)
	c.DELETE("/items/:item_id", ItemHandlerAPI.DeleteItem, middlewares.JWTMiddleware())
	c.GET("/items/:item_id", ItemHandlerAPI.GetItemByID)
	c.POST("/items", ItemHandlerAPI.CreateItem, middlewares.JWTMiddleware())
	c.PUT("/items/:item_id", ItemHandlerAPI.UpdateItemByID, middlewares.JWTMiddleware())
	c.GET("/tenantitems", ItemHandlerAPI.GetAllItemsTenant, middlewares.JWTMiddleware())
	c.GET("/archived", ItemHandlerAPI.SelectArchivedItem, middlewares.JWTMiddleware())
	c.PUT("/archive/:item_id", ItemHandlerAPI.Archive, middlewares.JWTMiddleware())
	c.PUT("/unarchive/:item_id", ItemHandlerAPI.Unarchive, middlewares.JWTMiddleware())

	//Admin
	c.POST("/admin", AdminHandlerAPI.Login)

	//Tenant
	c.POST("/tenant", TenantHandlerAPI.Insert)
	c.POST("/tenant/login", TenantHandlerAPI.Login)
	c.GET("/tenant", TenantHandlerAPI.GetAllTenant)
	c.GET("/tenant/:tenant_id/items", TenantHandlerAPI.GetTenantItems)
	c.DELETE("/tenant/:tenant_id", TenantHandlerAPI.DeleteTenant)
	c.GET("/tenant/:tenant_id", TenantHandlerAPI.GetTenantById)

	//Rent
	c.POST("/rent", RentHandlerAPI.CreateRent, middlewares.JWTMiddleware())
	c.GET("/rent/:rent_id", RentHandlerAPI.ReadRentById)
	c.PUT("/rent/:rent_id", RentHandlerAPI.UpdatebyId)
	c.POST("/rentpayment/:rent_id", RentHandlerAPI.Payment, middlewares.JWTMiddleware())
	c.POST("/callback", RentHandlerAPI.Callback)

	//Cart
	c.POST("/cart", CartHandlerAPI.CreateUserCart)
	c.GET("/cart/:cart_id", CartHandlerAPI.GetCartById)

	//Payment
	c.POST("/virtual-accounts", PaymentHandlerAPI.CreateVirtualAccount)
	c.GET("/virtual-accounts/:payment_id", PaymentHandlerAPI.GetVAById)
}
