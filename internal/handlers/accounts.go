package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EmailStruct struct {
	Email string
}

func Signup(c echo.Context) error {
	if c.Request().Method == "POST" {
		Email := c.Request().FormValue("email")
		// tools.SendEmail(email)
		emailData := EmailStruct{Email: Email}
		// emailData
		// fmt.Println(email)
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
	return nil
}
