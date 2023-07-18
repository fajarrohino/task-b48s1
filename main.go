package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, "hello my name is fajar rohino")
    })
    e.GET("/about", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string {
            "name" : "fajar rohino",
        })
    })
    e.Logger.Fatal(e.Start("localhost:5000"))
}