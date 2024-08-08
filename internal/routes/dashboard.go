package routes

import "github.com/labstack/echo/v4"

func InitDashboardRoutes(e *echo.Echo) {
	dashboardGroup := e.Group("/dashboard")
	dashboardGroup.GET("", nil)
}
