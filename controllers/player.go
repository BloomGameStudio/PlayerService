package controllers

import (
	"log"
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
)

func CreatePlayer(c echo.Context) error {

	// Creates a new Player in the Service
	// Expects a requestPlayer or a model.Player object in the body

	// Initializer request player to bind into
	reqPlayer := publicModels.Player{}

	err := c.Bind(&reqPlayer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	log.Printf("Bound the reqPlayer: %v", reqPlayer.Name)

	if !reqPlayer.IsValid() {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	playerModel := &models.Player{
		Position: reqPlayer.Position,
		Rotation: reqPlayer.Rotation,
		Scale:    reqPlayer.Scale,
	}

	if !playerModel.IsValid() {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	// save to db
	log.Println("Opening DB and saving playerModel")
	db := database.Open()
	db.Create(playerModel)

	return c.JSON(http.StatusCreated, playerModel)
}
