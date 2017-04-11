package views

import (
	"time"

	"github.com/labstack/echo"
)

// LoginView is show login form
func LoginView(c echo.Context, status int) error {
	dest := make(map[string]interface{})
	dest["Title"] = "Login"
	cookie, _ := c.Cookie("errormessage")
	if cookie != nil {
		dest["Error"] = cookie.Value
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		c.SetCookie(cookie)
	}

	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}

	return c.Render(status, "login", dest)
}

// LogoutView is show logout message
func LogoutView(c echo.Context, status int) error {
	dest := make(map[string]interface{})
	dest["Title"] = "Logout"
	deleteCookie(c, "uid")
	return c.Render(status, "logout", dest)
}
