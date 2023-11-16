package player

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func CreatePlayer(c echo.Context) error {

	// Creates a new Player in the Service
	// Expects a publicModel or a model dot Player object in the body

	c.Logger().Debug("In CreatePlayer")

	// Initializer request player to bind into
	reqPlayer := publicModels.Player{}

	err := c.Bind(&reqPlayer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	c.Logger().Debugf("Bound the player struct: %v", reqPlayer)

	if !reqPlayer.IsValid() {
		c.Logger().Debug("reqPlayer is NOT valid")
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	c.Logger().Debug("reqPlayer is valid")

	c.Logger().Debug("Initializing and populating player model!")
	// Use dot annotation for promoted aka embedded fields.
	playerModel := &models.Player{}

	// TODO: Set UserID
	// Note: TMP HOTFIX To create users in staging env without having to enter server or container
	if viper.GetBool("DEBUG") {
		// Sleep for 1 second in DEBUG mode to not get fludded with data
		playerModel.UserID = uuid.NewV4()
	}

	playerModel.Name = reqPlayer.Name
	playerModel.Position.Position = reqPlayer.Position
	playerModel.Rotation.Rotation = reqPlayer.Rotation
	playerModel.Scale.Scale = reqPlayer.Scale
	playerModel.States = []models.State{}
	for _, state := range reqPlayer.States {
		playerModel.States = append(playerModel.States, models.State{State: state})
	}
	//playerModel.ModelData.ModelID = reqPlayer.ModelData.ModelID
	playerModel.PlayerModel.MaterialID = reqPlayer.PlayerModel.MaterialID
	if !playerModel.IsValid() {
		c.Logger().Debug("playerModel is NOT valid")
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debug("playerModel is valid")
	db := database.GetDB()
	//Save the Player model
	if result := db.Create(&playerModel); result.Error != nil {
		c.Logger().Errorf("Failed to save playerModel: %v", result.Error)
		return c.JSON(http.StatusInternalServerError, "Failed to save player")
	}

	c.Logger().Debug("playerModel is saved. Returning")

	return c.JSON(http.StatusCreated, playerModel)
}
