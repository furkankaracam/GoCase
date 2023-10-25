package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Welcome(c echo.Context) error {
	return c.Render(http.StatusOK, "welcome", "")
}
