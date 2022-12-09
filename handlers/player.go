package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Player(player models.Player, c echo.Context) error {

	// c echo.Context is only used to obtain the logger to provide a unified logging
	logger := c.Logger()

	logger.Debug("We are in Player Handler")
	logger.Debug("player Arg: %v", player)

	db := database.Open()

	// Initialize empty database player model to bind into from db query
	databasePlayerModel := &models.Player{}

	var result *gorm.DB
	if viper.GetBool("DEBUG") {
		logger.Debug("Querying player by Name because we are in DEBUG mode")

		// Get the player by name in DEBUG mode for easier debugging and avoid the UserID dependency
		queryPlayer := &models.Player{}
		queryPlayer.Name = player.Name
		result = db.Preload(clause.Associations).Model(&models.Player{}).Where(queryPlayer).First(&databasePlayerModel)
	} else {
		logger.Debug("Querying database player by UserID")
		// Query db with the UserID from the passed in player model to find correct player
		result = db.Preload(clause.Associations).Model(&models.Player{}).Where(&models.Player{UserID: player.UserID}).First(&databasePlayerModel)
	}

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	logger.Debugf("Query result for databasePlayerModel %v", databasePlayerModel)
	// log.Print(helpers.PrettyStructNoError(databasePlayerModel))
	logger.Debug("Updating the databasePlayerModel")

	// It is assumed that the request can have a empty ID
	databasePlayerModel.Position = player.Position
	databasePlayerModel.Position.ID = databasePlayerModel.PositionID

	databasePlayerModel.Rotation = player.Rotation
	databasePlayerModel.Rotation.ID = databasePlayerModel.RotationID

	databasePlayerModel.Scale = player.Scale
	databasePlayerModel.Scale.ID = databasePlayerModel.ScaleID

	logger.Debugf("Updated databasePlayerModel: %v", databasePlayerModel)
	logger.Debug("Saving database player")

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&databasePlayerModel)
	// NOTE: it might be advantages to explicitly Save the updated fields like below to avoid accidental implicit updates:
	// db.Save(databasePlayerModel)
	// db.Save(databasePlayerModel.Position)
	logger.Debug("Returning")

	return nil

}
