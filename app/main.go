package main

import (
	"clean-arch-playground/message/delivery/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "/Users/nurayahmadova/WT/clean-arch-playground/app/static")

	h := http.NewHolder()
	e.GET("/ws", h.WSHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
