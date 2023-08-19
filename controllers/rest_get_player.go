package controllers

import (
	"log"
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

//Define a response struct for control over what gets serialized?

func GetPlayer(c echo.Context) error {
	// Open the database connection
	db := database.Open()
	//defer db.Close()
	if db == nil {
		log.Println("Failed to connect to the database")
		return c.JSON(http.StatusInternalServerError, "Failed to connect to the database")
	}

	queryPlayer := &models.Player{}
	queryPlayer.Active = true

	players := &[]models.Player{}
	if err := db.Preload(clause.Associations).Where(queryPlayer).Find(players).Error; err != nil {
		log.Println("Failed to retrieve players from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve players from the database")
	}

	// Return the list of players as a JSON response
	return c.JSON(http.StatusOK, players)
}
