package controllers

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
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
	playerModel.Name = reqPlayer.Name
	playerModel.Position = reqPlayer.Position
	playerModel.Rotation = reqPlayer.Rotation
	playerModel.Scale = reqPlayer.Scale

	if !playerModel.IsValid() {
		c.Logger().Debug("playerModel is NOT valid")
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debug("playerModel is valid")

	// Save to db
	c.Logger().Debug("Opening DB and saving playerModel")
	db := database.Open()
	db.Create(playerModel)

	c.Logger().Debug("playerModel is saved. Returning")

	return c.JSON(http.StatusCreated, playerModel)
}
