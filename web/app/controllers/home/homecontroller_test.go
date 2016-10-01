package homecontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/home", nil)
	if err != nil {
		t.Fatalf("echo GET /home req faild: %s", err)
	}
	rec := httptest.NewRecorder()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, Index(c)) {
		expected := "Hello, world"
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}
