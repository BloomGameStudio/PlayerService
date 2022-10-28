package main

import (
	"github.com/Balugrizzly/playerstate/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")

	// ws://localhost:1323/ws
	e.GET("/ws", controllers.Hello)

	e.Logger.Fatal(e.Start(":1323"))

}
