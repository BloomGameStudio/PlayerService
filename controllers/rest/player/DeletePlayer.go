package player

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"strconv"
    "regexp"
    "strings"
    uuid "github.com/satori/go.uuid"

)

func DeletePlayer(c echo.Context) error {
    db := database.GetDB()

    identifier := c.Param("identifier")
    parts := strings.Split(identifier, ":")

    queryPlayer := &models.Player{}

    uuidRegex, err := regexp.Compile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
    if err != nil {
    // Handle the error if the regex fails to compile
        return c.JSON(http.StatusInternalServerError, "Internal server error")
    }
    for _, part := range parts {
        // Use the compiled regex to match the string
        if uuidRegex.MatchString(part) {
            uuidValue, err := uuid.FromString(part)
            if err != nil {
                return c.JSON(http.StatusBadRequest, "Invalid UUID format")
            }
            queryPlayer.UserID = uuidValue
            continue
        }
        
        if idUint, err := strconv.ParseUint(part, 10, 64); err == nil {
            queryPlayer.ID = uint(idUint)
            continue
        }
        
        // If the part is not UUID or ID, then it's assumed to be a name.
        queryPlayer.Name = part
    }

    if len(parts) < 1 || len(parts) > 3 {
        return c.JSON(http.StatusBadRequest, "Invalid identifier format")
    }

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
