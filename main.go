package main

import (
	"Hannon-app/app/config"
	"Hannon-app/app/database"
	"Hannon-app/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitMysql(cfg)
	database.InittialMigration(dbMysql)

	// create a new echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMysql, e, cfg)
	//start server and port
	e.Logger.Fatal(e.Start(":444"))
}
