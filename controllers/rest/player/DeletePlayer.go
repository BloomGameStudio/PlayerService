package player

import (
	"net/http"
	"strconv"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func DeletePlayer(c echo.Context) error {
	db := database.GetDB()

	identifier := c.Param("id")
	var queryPlayer models.Player

	if uid, err := uuid.FromString(identifier); err == nil {
		queryPlayer.UserID = uid
	} else if id, err := strconv.Atoi(identifier); err == nil {
		queryPlayer.ID = uint(id)
	} else {
		queryPlayer.Name = identifier
	}

	if err := db.Where(&queryPlayer).First(&queryPlayer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(404, "Player not found")
		}
		// handle other database errors
		return c.JSON(500, "Database error")
	}

	result := db.Delete(&models.Player{}, queryPlayer.ID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete player")
	}
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "No player found to delete")
	}

	return c.JSON(http.StatusOK, "Player deleted successfully")
}
