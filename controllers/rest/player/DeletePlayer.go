package player

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"strconv"
	

)

func DeletePlayer(c echo.Context) error {
    // Open the database connection
    db := database.GetDB()

    idToRemove := c.Param("id")

    if idToRemove == "" {
        return c.JSON(http.StatusBadRequest, "Invalid 'id' parameter value. Use a valid ID.")
    }

    queryPlayer := &models.Player{}


	idUint, err := strconv.ParseUint(idToRemove, 10, 64)
	if err != nil {
		c.Logger().Error("Invalid 'id' parameter value. Use a valid numeric ID.")
		return c.JSON(http.StatusBadRequest, "Invalid 'id' parameter value. Use a valid numeric ID.")
	}
	queryPlayer.ID = uint(idUint)

    if err := db.Preload(clause.Associations).Where(queryPlayer).First(queryPlayer).Error; err != nil {
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