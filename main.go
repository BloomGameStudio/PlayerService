package main

import (
	"github.com/BloomGameStudio/PlayerService/controllers"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Enable Echo Debug mode if we are In DEBUG mode. Debug mode sets the log level to DEBUG.
	e.Debug = viper.GetBool("DEBUG")

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
	// ws.GET("position", controllers.Position)

	e.Logger.Fatal(e.Start(":1323"))

}
