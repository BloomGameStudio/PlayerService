package player

import (
	"net/http"

	"strconv"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

func UpdatePlayer(c echo.Context) error {

	db := database.GetDB()

	identifier := c.Param("id")

	if identifier == "" {
		return c.JSON(http.StatusBadRequest, "Invalid identifier parameter")
	}

	var queryPlayer models.Player

	if uid, err := uuid.FromString(identifier); err == nil {
		queryPlayer.UserID = uid
	} else if id, err := strconv.Atoi(identifier); err == nil {
		queryPlayer.ID = uint(id)
	} else {
		queryPlayer.Name = identifier
	}

	if err := db.Preload(clause.Associations).Where(&queryPlayer).First(&queryPlayer).Error; err != nil {
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

	queryPlayer.ModelData.MaterialID = updateData.ModelData.MaterialID

	// Save the updated player
	if err := db.Save(&queryPlayer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update player in the database")
	}

	// Return the updated player as a JSON response
	return c.JSON(http.StatusOK, queryPlayer)

}
