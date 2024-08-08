package main

import (
	"goinmemoryadmin/configs"
	"goinmemoryadmin/internal/routes"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {

	app := echo.New()

	app.Use(middleware.Logger())

	app.Renderer = configs.NewTemplate()

	routes.InitCoreRouter(app)
	routes.InitAccountsRoutes(app)

	app.Logger.Fatal(app.Start(":8080"))
}
