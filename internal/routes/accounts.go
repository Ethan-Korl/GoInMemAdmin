package routes

import (
	"goinmemoryadmin/internal/handlers"

	"github.com/labstack/echo/v4"
)

func InitAccountsRoutes(e *echo.Echo) {
	accountGroup := e.Group("/accounts")
	accountGroup.GET("/signup", handlers.Signup)
	accountGroup.POST("/signup", handlers.Signup)
	accountGroup.POST("/complete-signup", handlers.CompleteSignUp)
	accountGroup.POST("/login", handlers.Login)
}
