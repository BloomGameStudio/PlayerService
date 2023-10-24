package player

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"strconv"
    "strings"
	

)

func DeletePlayer(c echo.Context) error {
    // Open the database connection
    db := database.GetDB()

    identifier := c.Param("identifier")
    parts := strings.Split(identifier, ":")

    queryPlayer := &models.Player{}

    if len(parts) == 2 {
        if idUint, err := strconv.ParseUint(parts[0], 10, 64); err == nil {
            queryPlayer.ID = uint(idUint)
        }
        queryPlayer.Name = parts[1]
    } else {
        return c.JSON(http.StatusBadRequest, "Invalid identifier format")
    }

    // Delete the player and check if any rows were affected
    result := db.Where(queryPlayer).Delete(&models.Player{})
    if result.Error != nil {
        c.Logger().Error("Failed to delete player from the database")
        return c.JSON(http.StatusInternalServerError, "Failed to delete player from the database")
    }
    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "No player found to delete")
    }
 
    return c.JSON(http.StatusOK, "Player deleted successfully")
}
