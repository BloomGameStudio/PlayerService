package helpers

import (
	"bytes"
	"encoding/json"

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

	db := database.GetDB()
	result := db.Model(&models.Player{}).Where(&models.Player{UserID: userID}).First(&playerModel)

	if result.Error != nil {
		return nil, result.Error
	}

	return playerModel, nil
}

func PrettyStruct(data interface{}) (string, error) {
	// Beautifies a struct
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func PrettyStructNoError(data interface{}) string {
	// Beautifies a struct returns empty string on error
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return ""
	}
	return string(val)
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
