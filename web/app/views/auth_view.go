package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// LoginView is root
func LoginView(c echo.Context) error {
	dest := make(map[string]interface{})
	cookie, _ := c.Cookie("errormessage")
	if cookie != nil {
		dest["Error"] = cookie.Value
		fmt.Println(cookie.Value)
		cookie.Value = ""
		cookie.Expires = time.Now()
		c.SetCookie(cookie)
	}

	return c.Render(http.StatusOK, "login", dest)
}
