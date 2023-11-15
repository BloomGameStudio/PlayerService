package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Level(level models.Level, c echo.Context) error {

	// c echo.Context is only used to obtain the logger to provide unified logging
	logger := c.Logger()

	db := database.GetDB()

	// Initialize empty database model to bind into from db query
	databaseLevelModel := &models.Level{}

	var result *gorm.DB

	// Query db with the ID from the passed in model

	queryLevel := &models.Level{}
	queryLevel.ID = level.ID

	result = db.Model(&models.Level{}).Where(queryLevel).FirstOrCreate(&databaseLevelModel, &queryLevel)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	databaseLevelModel.LevelID = level.LevelID

	// NOTE: This is for debug mode reevaluate PlayerID
	if viper.GetBool("DEBUG") {
		// Accept client provided PlayerID in DEBUG mode
		databaseLevelModel.PlayerID = level.PlayerID

	}

	db.Updates(&databaseLevelModel)

	return nil

}
