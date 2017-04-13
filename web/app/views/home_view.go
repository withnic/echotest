package views

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/withnic/echotest/web/app/models"
)

// HomeView is root
func HomeView(c echo.Context, messages []models.MessageWithUser) error {
	dest := make(map[string]interface{})
	dest["Title"] = "Home"
	dest["MessageWithUser"] = messages

	current, err := current(c)
	if err == nil {
		dest["Current"] = current
	}

	return c.Render(http.StatusOK, "homeindex", dest)
}
