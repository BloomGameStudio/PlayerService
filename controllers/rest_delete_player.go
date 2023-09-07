package controllers

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func DeletePlayer(c echo.Context) error {
	//Open the database connection
	db := database.GetDB()

	idToRemove := c.QueryParam("id")

	if idToRemove == "" {
		return c.JSON(http.StatusBadRequest, "Invalid 'id' parameter value. Use a valid ID.")
	}

	queryPlayer := &models.Player{}
	if err := db.Preload(clause.Associations).Where("id = ?", idToRemove).First(queryPlayer).Error; err != nil {
		c.Logger().Error("Failed to retrieve player from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve player from the database")
	}


	// Delete the player
	if err := db.Delete(queryPlayer).Error; err != nil {
		c.Logger().Error("Failed to delete player from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to delete player from the database")
	}

	return c.JSON(http.StatusOK, "Player deleted successfully")
}