package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Position(position models.Position, c echo.Context) error {

	// c echo.Context is only used to obtain the logger to provide unified logging
	logger := c.Logger()

	db := database.GetDB()

	// Initialize empty database position model to bind into from db query
	databasePositionModel := &models.Position{}

	var result *gorm.DB

	logger.Debug("Querying database position by ID")
	// Query db with the ID from the passed in position model to find correct player

	queryPosition := &models.Position{}
	queryPosition.ID = position.ID

	result = db.Model(&models.Position{}).Where(queryPosition).First(&databasePositionModel)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	databasePositionModel.Position = position.Position

	db.Updates(&databasePositionModel)

	return nil

}
