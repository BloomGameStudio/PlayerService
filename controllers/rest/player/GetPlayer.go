package player

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

//Define a response struct for control over what gets serialized?

func GetPlayer(c echo.Context) error {
	// Open the database connection
	db := database.GetDB()

	// Read the "active" query parameter from the URL
	activeParam := c.QueryParam("active")

	// Initialize a variable to store the filter value
	var active bool

	// Check if the "active" query parameter is provided and parse it as a boolean
	switch activeParam {
	case "true":
		active = true
	case "false":
		active = false
	case "":
		active = true

	default:
		return c.JSON(http.StatusBadRequest, "Invalid 'active' parameter value. Use 'true' or 'false'.")
	}

	// Build the query based on the "active" filter
	queryPlayer := &models.Player{}
	queryPlayer.Active = active
	players := &[]models.Player{}
	if err := db.Preload(clause.Associations).Where(queryPlayer).Find(players).Error; err != nil {
		c.Logger().Error("Failed to retrieve players from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve players from the database")
	}
	// Return the list of players as a JSON response
	return c.JSON(http.StatusOK, players)
}