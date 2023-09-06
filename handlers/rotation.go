package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Rotation(rotation models.Rotation, c echo.Context) error {

	// c echo.Context is only used to obtain the logger to provide unified logging
	logger := c.Logger()

	logger.Debug("We are in Rotation Handler")
	logger.Debug("rotation Arg: %v", rotation)

	db := database.GetDB()

	// Initialize empty database rotation model to bind into from db query
	databaseRotationModel := &models.Rotation{}

	var result *gorm.DB

	logger.Debug("Querying database rotation by ID")
	// Query db with the ID from the passed in rotation model to find correct player

	queryRotation := &models.Rotation{}
	queryRotation.ID = rotation.ID

	result = db.Model(&models.Rotation{}).Where(queryRotation).First(&databaseRotationModel)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	logger.Debugf("Query result for databaseRotationModel %v", databaseRotationModel)
	// log.Print(helpers.PrettyStructNoError(databasePlayerModel))
	logger.Debug("Updating the databaseRotationModel")

	databaseRotationModel.Rotation = rotation.Rotation

	logger.Debugf("Updated databaseRotationModel: %v", databaseRotationModel)
	logger.Debug("Saving database Rotation")

	db.Updates(&databaseRotationModel)

	return nil

}
