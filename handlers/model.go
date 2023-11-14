package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
func ModelHandler(modelData models.Model, c echo.Context) error {
	logger := c.Logger()

	logger.Debug("We are in ModelHandler")
	logger.Debugf("modelData Arg: %v", modelData)

	db := database.GetDB()
	databaseModelData := &models.Model{}

	var result *gorm.DB

	logger.Debugf("Querying database model by ID: %d", modelData.ID)
	result = db.Model(&models.Model{}).Where("id = ?", modelData.ID).First(&databaseModelData)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	logger.Debugf("Query result for databaseModelData %v", databaseModelData)
	logger.Debug("Updating the databaseModelData")

	// Update the values of databaseModelData
	databaseModelData.ModelID = modelData.ModelID
	databaseModelData.MaterialID = modelData.MaterialID

	logger.Debugf("Updated databaseModelData: %v", databaseModelData)
	logger.Debug("Saving database Model")

	db.Save(&databaseModelData)

	return nil
}