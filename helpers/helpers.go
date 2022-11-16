package helpers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func GetPlayerModelFromJWT(c echo.Context) (playerModel *models.Player, err error) {

	jwtClaims := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)

	// Use more verbose and lengthy db lookup query for safety instead of:
	// result := db.Model(&models.User{}).Where("ID = ?", userID).First(&userModel)
	userID, err := uuid.FromString(jwtClaims["ID"].(string))

	if err != nil {
		panic(err)
	}

	db := database.Open()
	result := db.Model(&models.Player{}).Where(&models.Player{UserID: userID}).First(&playerModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return playerModel, nil
}
