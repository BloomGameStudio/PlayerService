package controllers

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
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
	// db := database.Open()

	// OPTIMIZE: Use GetUserIDFromJWT function to avoid db call
	// TODO: UNCOMNNET
	// player, err := helpers.GetPlayerModelFromJWT(c)

	// if err != nil {
	// 	return err
	// }
	// playerUserID := player.UserID

	for {

		// Write
		// func() {

		// 	// Get all the players
		// 	players := &models.Player{}
		// 	db.Find(players)

		// 	// Find/Filter the Changes that occured in the players and send them
		// 	// PlayerChanges(players,players)

		// 	err := ws.WriteJSON(players)
		// 	if err != nil {
		// 		c.Logger().Error(err)
		// 	}
		// }()

		// Read
		func() {

			// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!

			// Initializer request player to bind into
			reqPlayer := &publicModels.Player{}

			err := ws.ReadJSON(reqPlayer)

			pretyReqPlayer, _ := PrettyStruct(reqPlayer)

			log.Printf("reqPlayer: %+v", pretyReqPlayer)

			if err != nil {

				log.Printf("We get an error from Reading the JSON reqPlayer")
				c.Logger().Error(err)
			}

			playerModel := &models.Player{
				// UserID:   playerUserID,
				Position: reqPlayer.Position,
				Rotation: reqPlayer.Rotation,
				Scale:    reqPlayer.Scale,
			}

			if viper.GetBool("DEBUG") {
				// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
				playerModel.Name = reqPlayer.Name
			}

			prettyPlayerModel, _ := PrettyStruct(playerModel)
			log.Printf("playerModel: %+v", prettyPlayerModel)

			handlers.Player(*playerModel)

		}()

	}
}
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
