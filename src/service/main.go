package main

import (
	"github.com/ChichaSasha/BitMediaTest/service/api"
	"github.com/labstack/echo"
	"net/http"

	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	m := api.Manager("1")

	api.Assemble(e, m)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
