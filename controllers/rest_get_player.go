package controllers

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

//Define a response struct for contfffrol over what gets serialized?

func GetPlayer(c echo.Context) error {
	// Open the database connection
	db := database.GetDB()

	queryPlayer := &models.Player{}
	queryPlayer.Active = true

	players := &[]models.Player{}
	if err := db.Preload(clause.Associations).Where(queryPlayer).Find(players).Error; err != nil {
		c.Logger().Error("Failed to retrieve players from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve players from the database")
	}

	// Return the list of players as a JSON response
	return c.JSON(http.StatusOK, players)
}
