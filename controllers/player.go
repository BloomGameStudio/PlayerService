package controllers

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

// TODO: implement requestPlayer
type requestPlayer struct {
	Placeholder string `json:"name"`
}

func CreatePlayer(c echo.Context) error {

	// Creates a new Player in the Service
	// Expects a requestPlayer or a model.Player object in the body

	// Initializer request player to bind into
	reqPlayer := requestPlayer{}

	err := c.Bind(&reqPlayer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	if requestPlayer.isValid() != true {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	playerModel := &models.Player{
		Name: reqPlayer.Placeholder,
	}

	if playerModel.isValid() != true {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	// save to db
	db := database.Open()
	db.Create(playerModel)

	return c.JSON(http.StatusCreated, playerModel)
}
