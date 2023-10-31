package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ModelHandler(modelData models.Model, c echo.Context) error {

	// c echo.Context is only used to obtain the logger to provide unified logging
	logger := c.Logger()

	logger.Debug("We are in ModelHandler")
	logger.Debug("modelData Arg: %v", modelData)

	db := database.GetDB()

	// Initialize an empty database model to bind into from db query
	databaseModelData := &models.Model{}

	var result *gorm.DB

	logger.Debug("Querying database model by ID")
	// Query db with the ID from the passed in model to find the correct model

	queryModel := &models.Model{}
	queryModel.ID = modelData.ID

	result = db.Model(&models.Model{}).Where(queryModel).First(&databaseModelData)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	logger.Debugf("Query result for databaseModelData %v", databaseModelData)
	logger.Debug("Updating the databaseModelData")

	// Update the ModelData field
	databaseModelData.ModelData = modelData.ModelData

	logger.Debugf("Updated databaseModelData: %v", databaseModelData)
	logger.Debug("Saving database Model")

	db.Updates(&databaseModelData)

	return nil
}
