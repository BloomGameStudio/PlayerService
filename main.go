package main

import (
	"github.com/BloomGameStudio/PlayerService/controllers/rest/ping"
	"github.com/BloomGameStudio/PlayerService/controllers/rest/player"
	"github.com/BloomGameStudio/PlayerService/controllers/rest/version"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/hello"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/level"
	wsPlayer "github.com/BloomGameStudio/PlayerService/controllers/ws/player"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/position"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/rotation"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/scale"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/state"
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
	e.GET("ping", ping.Ping)
	e.GET("v", version.Version)
	// End of HTTP testing routes

	// Player routes
	e.GET("player", player.GetPlayer)
	e.POST("player", player.CreatePlayer)
	e.PUT("player/:identifier", player.UpdatePlayer)
	// End Player routes

	// End of HTTP routes

	// WebSocket Routes
	ws := e.Group("/ws/")

	// Web Socket Testing routes
	ws.File("", "public/index.html") // http://127.0.0.1:1323/ws/
	// ws://localhost:1323/ws
	ws.GET("hello", hello.Hello)
	// End of Web Socket testing routes

	ws.GET("player", wsPlayer.Player)
	ws.GET("state", state.State)
	ws.GET("position", position.Position)
	ws.GET("rotation", rotation.Rotation)
	ws.GET("scale", scale.Scale)
	ws.GET("level", level.Level)

	port := viper.GetString("PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
