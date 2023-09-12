package player

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func playerReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {

		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the WebSocket")

			// Initializer request player to bind into
			reqPlayer := &publicModels.Player{}
			err := ws.ReadJSON(reqPlayer)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqPlayer")
				switch {

				case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
					c.Logger().Debug("Websocket CloseNoStatusReceived")
					select {

					case ch <- nil:
						c.Logger().Debug("Sent nil to Reader channel")
						return

					case <-time.After(time.Second * 10):
						c.Logger().Debug("Timed out sending nil to Reader channel")
						return
					}

				default:

					c.Logger().Error(err)

					select {

					case ch <- err:
						c.Logger().Debug("Sent error to Reader channel")
						return
					case <-time.After(time.Second * 10):
						c.Logger().Debug("Timed out sending error to Reader channel")
						return
					}
				}
			}

			c.Logger().Debugf("reqPlayer from the WebSocket: %+v", reqPlayer)

			c.Logger().Debug("Validating reqPlayer")
			if !reqPlayer.IsValid() {

				c.Logger().Debug("reqPlayer is NOT valid returning")
				ch <- errors.New("reqPlayer Validation failed")
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return
			}

			c.Logger().Debug("reqPlayer is valid")

			c.Logger().Debug("Initializing and populating player model!")
			// Use dot annotation for promoted aka embedded fields.
			playerModel := &models.Player{}
			// TODO: Handle UserID and production mode
			playerModel.Position.Position = reqPlayer.Position
			playerModel.Rotation.Rotation = reqPlayer.Rotation
			playerModel.Scale.Scale = reqPlayer.Scale

			for _, state := range reqPlayer.States {
				playerModel.States = append(playerModel.States, models.State{State: state})
			}

			playerModel.Layer = reqPlayer.Layer
			playerModel.PModel = reqPlayer.PModel
			playerModel.Active = reqPlayer.Active

			if viper.GetBool("DEBUG") {
				// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
				playerModel.Name = reqPlayer.Name
			}

			c.Logger().Debugf("playerModel: %+v", playerModel)

			c.Logger().Debug("Validating playerModel")
			if !playerModel.IsValid() {

				c.Logger().Debug("playerModel is NOT valid returning")
				// NOTE: No Timeout used here
				ch <- errors.New("playerModel Validation failed")
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return
			}

			c.Logger().Debug("playerModel is valid passing it to the Player handler")
			handlers.Player(*playerModel, c) //TODO: handle errors
		}
	}

}
