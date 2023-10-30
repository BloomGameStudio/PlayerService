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

    identifier := c.Param("identifier")
    var query string
    var arg interface{}

    if _, err := uuid.FromString(identifier); err == nil {
        query, arg = "user_id = ?", identifier
    } else if _, err := strconv.Atoi(identifier); err == nil {
        id, _ := strconv.Atoi(identifier)
        query, arg = "id = ?", id
    } else {
        query, arg = "name = ?", identifier
    }
    var queryPlayer models.Player
	if err := db.Where(query, arg).First(&queryPlayer).Error; err != nil {
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
