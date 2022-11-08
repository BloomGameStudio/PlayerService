package main

import (
	"github.com/BloomGameStudio/PlayerService/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Testing routes
	e.GET("ping", controllers.Ping)
	e.File("/", "public/index.html")
	// ws://localhost:1323/ws
	e.GET("/ws", controllers.Hello)
	// End of testing routes

	e.POST("player", controllers.CreatePlayer)

	e.Logger.Fatal(e.Start(":1323"))

}
