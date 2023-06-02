package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var i = echo.New()

	i.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	i.Logger.Fatal(i.Start("localhost:5000"))

}
