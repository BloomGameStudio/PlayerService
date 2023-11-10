package config

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/spf13/viper"
)

func Init() {
	ViperInit()
	// Initialize Database
	database.Init()

	// TODO: Cleanup Migration
	db := database.GetDB()

	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.State{})
	db.AutoMigrate(&models.Model{})
	db.AutoMigrate(&models.Level{})

}

func ViperInit() {
	// Initialize Viper with project configuration.

	// Setting Default Values
	viper.SetDefault("DEBUG", true)
	viper.SetDefault("PORT", "1323") // Has to be a string as Echo expects a string
	viper.SetDefault("WS_TIMEOUT_SECONDS", 10)
}
