package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PlayerModel(playerModel models.PlayerModel, c echo.Context) error {
	logger := c.Logger()

	logger.Debug("We are in ModelHandler")
	logger.Debugf("modelData Arg: %v", playerModel)

	db := database.GetDB()
	databaseModelData := &models.PlayerModel{}

	var result *gorm.DB

	logger.Debugf("Querying database model by ID: %d", playerModel.ID)
	result = db.Model(&models.PlayerModel{}).Where(models.PlayerModel{PlayerID: playerModel.ID}).FirstOrCreate(&databaseModelData)

	if result.Error != nil {
		logger.Debugf("Error querying or creating playerModel for ID: %d", playerModel.ID)
		return result.Error
	}

	logger.Debugf("Query result for databaseModelData %v", databaseModelData)
	logger.Debug("Updating the databaseModelData")

	// Update the values of databaseModelData
	databaseModelData.MaterialID = playerModel.MaterialID

	logger.Debugf("Updated databaseModelData: %v", databaseModelData)
	logger.Debug("Saving database Model")

	db.Save(&databaseModelData)

	return nil
}
