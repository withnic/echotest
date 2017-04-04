package home

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"github.com/withnic/echotest/web/app/libs"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	libs.SetRenderer(e, "../../template/**/*.html")
	req, err := http.NewRequest(echo.GET, "/home", nil)
	if err != nil {
		t.Fatalf("echo GET /home req faild: %s", err)
	}
	rec := httptest.NewRecorder()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, Index(c)) {
		expected := "<title>Hello, world</title>"
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Regexp(t, expected, rec.Body.String(), "dose not has title")
	}
}
