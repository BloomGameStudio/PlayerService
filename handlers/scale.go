package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

func Scale(scale models.Scale, c echo.Context) error {
	db := database.GetDB()

	databaseScaleModel := &models.Scale{}

	queryScale := &models.Scale{}
	queryScale.ID = scale.ID

	result := db.Model(&models.Scale{}).Where(queryScale).First(&databaseScaleModel)

	if result.Error != nil {
		return result.Error
	}

	databaseScaleModel.Scale = scale.Scale

	db.Updates(&databaseScaleModel)

	return nil
}
