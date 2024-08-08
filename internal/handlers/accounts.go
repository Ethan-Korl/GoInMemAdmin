package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	if c.Request().Method == "POST" {
		email := c.Request().FormValue("email")
		// tools.SendEmail(email)
		// time.Sleep(time.Duration(12334833939338849))
		var emailData = struct{ Email string }{Email: email}
		return c.Render(http.StatusOK, "confirm_email.html", emailData)
	}

	fmt.Println("GET METHOD")
	return c.Render(http.StatusOK, "signup.html", nil)
}

func CompleteSignUp(c echo.Context) error {
	username := c.Request().FormValue("username")
	password := c.Request().FormValue("password")

	fmt.Println(username, password)
	return c.Render(http.StatusOK, "signup.html", nil)
}

func Login(c echo.Context) error {
	username := c.Request().FormValue("username")
	password := c.Request().FormValue("password")
	fmt.Println(username, password)
	return c.Render(http.StatusOK, "signup.html", nil)
}
