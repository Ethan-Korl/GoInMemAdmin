package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	// tmpl := template.Must(template.ParseFiles("index.html"))
	fmt.Println(os.Getenv("name"))
	return c.Render(http.StatusOK, "index.html", nil)
}
