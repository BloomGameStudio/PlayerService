package controllers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

func Player(c echo.Context) error {

	// QUESTION: Is this needed?
	// Only changes will be sent the only exception to this is the opening/first request where the full state will be sent
	// Partial player data can be received or full
	// TODO: Partial Reads
	// TODO: Partial Writes

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Open DB outside of loopception
	db := database.Open()

	for {

		// Write
		func() {

			// Get all the players
			players := &models.Player{}
			db.Find(players)

			// Find/Filter the Changes that occured in the players and send them
			// PlayerChanges(players,players)

			err := ws.WriteJSON(players)
			if err != nil {
				c.Logger().Error(err)
			}
		}()

		// Read
		func() {

			// Initializer request player to bind into
			reqPlayer := &RequestPlayer{}

			err := ws.ReadJSON(reqPlayer)

			if err != nil {
				c.Logger().Error(err)
			}

			handlers.Player(reqPlayer)

		}()

	}
}