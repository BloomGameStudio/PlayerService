package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
)

func Player(player models.Player) bool {

	db := database.Open()

	// Initialize empty database player model to bind into from db query
	databasePlayerModel := &models.Player{}

	// Query db with the UserID from the passed in player model to find correct player
	result := db.Model(&models.Player{}).Where(&models.Player{UserID: player.UserID}).First(&databasePlayerModel)

	if result.Error != nil {
		// TODO: handle error
		panic(result.Error)
	}

	// TODO: Check Updating Functionality and behaviour(also in regards to partial updates)
	// Updating the database player model with the new data from player argument
	databasePlayerModel.Position = player.Position
	databasePlayerModel.Rotation = player.Rotation
	databasePlayerModel.Scale = player.Scale

	db.Save(databasePlayerModel)

	return true

}
