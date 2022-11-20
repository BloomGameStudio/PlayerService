package controllers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
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

	// OPTIMIZE: Use GetUserIDFromJWT function to avoid db call
	// TODO: UNCOMNNET
	// player, err := helpers.GetPlayerModelFromJWT(c)

	// if err != nil {
	// 	return err
	// }
	// playerUserID := player.UserID

	for {

		// TODO: Close connection handling

		// Write
		func() {

			c.Logger().Debug("Writing to the WebSocket")

			c.Logger().Debug("Getting all players from the database")
			// Get all the players
			players := &models.Player{}
			db.Preload(clause.Associations).Find(players)

			// Find/Filter the Changes that occured in the players and send them
			// PlayerChanges(players,players)

			c.Logger().Debug("Pushing the player to the WebSocket")
			err := ws.WriteJSON(players)
			if err != nil {
				c.Logger().Error(err)
			}
			c.Logger().Debug("Finished writing to the WebSocket")

			// time.Sleep(time.Second * 3)

		}()

		// Read
		func() {

			// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
			// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

			c.Logger().Debug("Reading from the WebSocket")

			// Initializer request player to bind into
			reqPlayer := &publicModels.Player{}
			err := ws.ReadJSON(reqPlayer)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqPlayer")
				c.Logger().Error(err)
			}

			c.Logger().Debugf("reqPlayer from the WebSocket: %+v", reqPlayer)

			c.Logger().Debug("Validating reqPlayer")
			if !reqPlayer.IsValid() {
				c.Logger().Debug("reqPlayer is NOT valid returning")
				return
			}

			c.Logger().Debug("reqPlayer is valid")

			c.Logger().Debug("Initializing and populating player model!")
			// Use dot annotation for promoted aka embedded fields.
			playerModel := &models.Player{}
			// TODO: Handle UserID and production mode
			playerModel.Position = reqPlayer.Position
			playerModel.Rotation = reqPlayer.Rotation
			playerModel.Scale = reqPlayer.Scale

			if viper.GetBool("DEBUG") {
				// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
				playerModel.Name = reqPlayer.Name
			}

			c.Logger().Debugf("playerModel: %+v", playerModel)

			c.Logger().Debug("Validating playerModel")
			if !playerModel.IsValid() {
				c.Logger().Debug("playerModel is NOT valid returning")
				return
			}

			c.Logger().Debug("playerModel is valid passing it to the Player handler")
			// handlers.Player(*playerModel) //TODO: UNCOMNNET

		}()

	}
}
