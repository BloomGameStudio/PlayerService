package player

import (
	"net/http"

	"strconv"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func UpdatePlayer(c echo.Context) error {

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
	updateData := models.Player{}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid update data")
	}

	// Update the specific fields in queryPlayer
	queryPlayer.Layer = updateData.Layer
	queryPlayer.ENS = updateData.ENS
	queryPlayer.Active = updateData.Active
	queryPlayer.Transform.Position = updateData.Transform.Position
	queryPlayer.Transform.Rotation = updateData.Transform.Rotation
	queryPlayer.Transform.Scale = updateData.Transform.Scale

	queryPlayer.States = []models.State{}

	for _, state := range updateData.States {
		queryPlayer.States = append(queryPlayer.States, models.State{
			State: state.State,
		})
	}

	// Save the updated player
	if err := db.Save(queryPlayer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update player in the database")
	}

	// Return the updated player as a JSON response
	return c.JSON(http.StatusOK, queryPlayer)

}
