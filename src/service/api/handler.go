package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Assemble(e *echo.Echo, m Manager) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/users")

	g.GET("", h.users)
	g.GET("/rating", h.rating)
}

type handler struct {
	manager Manager
}

// handlers

func (h *handler) users(c echo.Context) error {

	return c.JSON(http.StatusCreated, map[string]int{"id": 1})
}

func (h *handler) rating(c echo.Context) error {
	return fmt.Errorf("not implemented")
}