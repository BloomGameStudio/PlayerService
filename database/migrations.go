package database

import "github.com/BloomGameStudio/PlayerService/models"

func Migrate() {

	db := Open()
	// This will Auto Migrate all its nested structs
	db.AutoMigrate(&models.Health{})
}
