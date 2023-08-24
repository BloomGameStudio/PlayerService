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
	// This will Auto Migrate all its nested structs
	db.AutoMigrate(&models.Player{})

}

func ViperInit() {
	// Initialize Viper with project configuration.

	// Setting Default Values
	viper.SetDefault("DEBUG", true)
	viper.SetDefault("PORT", "1323") // Has to be a string as Echo expects a string
}
