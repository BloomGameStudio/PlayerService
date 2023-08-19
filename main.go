package main

import (
	"github.com/BloomGameStudio/PlayerService/controllers"
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

	// HTTP routes

	// HTTP Testing routes
	e.GET("ping", controllers.Ping)
	e.GET("v", controllers.Version)
	e.GET("getplayer", controllers.GetPlayer)
	// End of HTTP testing routes

	e.POST("player", controllers.CreatePlayer)
	// End of HTTP routes

	// WebSocket Routes
	ws := e.Group("/ws/")

	// Web Socket Testing routes
	ws.File("", "public/index.html") // http://127.0.0.1:1323/ws/
	// ws://localhost:1323/ws
	ws.GET("hello", controllers.Hello)
	// End of Web Socket testing routes

	ws.GET("player", controllers.Player)
	ws.GET("state", controllers.State)
	ws.GET("position", controllers.Position)

	port := viper.GetString("PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
