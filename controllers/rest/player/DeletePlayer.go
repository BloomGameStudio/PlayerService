package player

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
    uuid "github.com/satori/go.uuid"
	"strconv"
	

)

func DeletePlayer(c echo.Context) error {
    // Open the database connection
    db := database.GetDB()

    idParam := c.QueryParam("id")
    nameParam := c.QueryParam("name")
    userIDParam := c.QueryParam("userid")

    queryPlayer := &models.Player{}

    //check ID
    if idParam != "" {
        idUint, err := strconv.ParseUint(idParam, 10, 64)
        if err != nil {
            c.Logger().Error("Invalid 'id' parameter value. Use a valid numeric ID.")
            return c.JSON(http.StatusBadRequest, "Invalid 'id' parameter value. Use a valid numeric ID.")
        }
        queryPlayer.ID = uint(idUint)
    }

    // Check Name
    if nameParam != "" {
        queryPlayer.Name = nameParam
    }

    // Check UserID
    if userIDParam != "" {
        userID, err := uuid.FromString(userIDParam)
        if err != nil {
            c.Logger().Error("Invalid 'userid' parameter value. Use a valid UUID.")
            return c.JSON(http.StatusBadRequest, "Invalid 'userid' parameter value. Use a valid UUID.")
        }
        queryPlayer.UserID = userID
    }
    

    // Make sure at least one parameter is provided
    if idParam == "" && nameParam == "" && userIDParam == "" {
        return c.JSON(http.StatusBadRequest, "Provide at least one parameter to delete a player")
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