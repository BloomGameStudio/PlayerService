package controllers

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

// OPTIMIZE: Staticly link requestPlayer with Model.Player
type requestPlayer struct {
	// COMBAK: Add needed further fields from the Player struct model
	Position models.Position `json:"position"`
	Rotation models.Rotation `json:"rotation"`
	Scale    models.Scale    `json:"scale"`
}

func (rp requestPlayer) isValid() bool {

	// Validates the requestPlayer
	// Additional validation and hooks for the reqeurequestPlayer validation can be added here
	// WARNING: Validation should be scoped to the requestPlayer

	if !rp.Position.IsValid() {
		return false
	}

	if !rp.Rotation.IsValid() {
		return false
	}

	if !rp.Scale.IsValid() {
		return false
	}

	return true
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

	if !reqPlayer.isValid() {
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
	db := database.Open()
	db.Create(playerModel)

	return c.JSON(http.StatusCreated, playerModel)
}
