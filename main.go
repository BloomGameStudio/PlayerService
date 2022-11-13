package main

import (
	"github.com/BloomGameStudio/PlayerService/controllers"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := database.Open()
	// This will Auto Migrate all its nested structs
	db.AutoMigrate(&models.Player{})

	// HTTP routes

	// HTTP Testing routes
	e.GET("ping", controllers.Ping)
	// End of HTTP testing routes

	e.POST("player", controllers.CreatePlayer)
	// End of HTTP routes

	// WebSocket Routes
	ws := e.Group("/ws/")

	// Web Socket Testing routes
	ws.File("", "public/index.html") // http://127.0.0.1:1323/ws/
	// ws://localhost:1323/ws
	ws.GET("hello", controllers.Hello)
	// End of Web Socket esting routes

	ws.GET("player", controllers.Player)

	e.Logger.Fatal(e.Start(":1323"))

}
