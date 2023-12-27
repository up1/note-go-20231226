package demo_test

import (
	"demo"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestUnitHelloAPI(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	// Assertions
	expect := `{"message_123":"Hello, World!"}`
	if assert.NoError(t, demo.Hello(&MockHelloRepo{})(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expect, strings.TrimSpace(rec.Body.String()))
	}
}

type MockHelloRepo struct{}

func (r *MockHelloRepo) GetData() (string, error) {
	return "Hello, World!", nil
}

func TestUnitPaicAPI(t *testing.T) {
	e := echo.New()
	e.Use(middleware.Recover())
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	// Assertions
	expect := `{"message":"Internal Server Error"}`
	if assert.NoError(t, demo.TryToFail(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expect, strings.TrimSpace(rec.Body.String()))
	}
}
