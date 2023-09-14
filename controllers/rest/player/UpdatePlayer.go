package player

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"strconv"
)
func UpdatePlayer(c echo.Context) error {

	//we send the entire player struct to the client
	
	// Open the database connection
	db := database.GetDB()

	// Parameters
	playerIDStr := c.Param("id")

	if playerIDStr == "" {
    	return c.JSON(http.StatusBadRequest, "Invalid id parameter value. Use a valid ID")
	}

	// Convert playerIDStr to uint
	playerID, err := strconv.ParseUint(playerIDStr, 10, 64)
	if err != nil {
    	return c.JSON(http.StatusBadRequest, "Invalid id parameter value. Use a valid ID")
	}

	// Find the player from ID given
	queryPlayer := &models.Player{}
	if err := db.Preload(clause.Associations).Where(playerID).First(queryPlayer).Error; err != nil {
    	return c.JSON(http.StatusNotFound, "Failed to retrieve player from the database")
	}

	// Parse the JSON request body into a separate variable
	updateData := struct {
		Layer string `json:"layer"`
		Position models.Position `json:"position"`
		Rotation models.Rotation `json:"rotation"`
		Scale    models.Scale    `json:"scale"`
	}{}

	if err := c.Bind(&updateData); err != nil {
		c.Logger().Error("Failed to parse update data")
		return c.JSON(http.StatusBadRequest, "Invalid update data")
	}

	// Update the specific fields in queryPlayer
	queryPlayer.Position = updateData.Position
	queryPlayer.Rotation = updateData.Rotation
	queryPlayer.Scale = updateData.Scale

	// Save the updated player
	if err := db.Save(queryPlayer).Error; err != nil {
		c.Logger().Error("Failed to update player in the database")
		return c.JSON(http.StatusInternalServerError, "Failed to update player in the database")
	}

	// Return the updated player as a JSON response
	return c.JSON(http.StatusOK, queryPlayer)
}

