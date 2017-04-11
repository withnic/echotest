package views

import (
	"net/http"

	"github.com/labstack/echo"
)

// MessagesFormView is Message create Form Page Helper
func MessagesFormView(c echo.Context) error {
	dest := make(map[string]interface{})
	dest["Title"] = "Message"

	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}
	return c.Render(http.StatusOK, "messagenew", dest)
}
