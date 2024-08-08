package routes

import (
	"goinmemoryadmin/internal/handlers"

	"github.com/labstack/echo/v4"
)

func InitCoreRouter(e *echo.Echo) {
	coreGroup := e.Group("/")
	coreGroup.GET("", handlers.Index)
}
