package player

import (
	"context"
	"errors"

	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
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
			reqPlayerArr := &[]publicModels.Player{}
			err := ws.ReadJSON(reqPlayerArr)

			if err != nil {
				errorHandlers.HandleReadError(c, ch, err, false)
				return
			}

			for _, reqPlayer := range *reqPlayerArr {

				c.Logger().Debugf("reqPlayer from the WebSocket: %+v", reqPlayer)

				if !reqPlayer.IsValid() {

					ch <- errors.New("reqPlayer Validation failed")
					return
				}

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
				playerModel.Active = reqPlayer.Active

				if viper.GetBool("DEBUG") {
					// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
					playerModel.Name = reqPlayer.Name
				}

				c.Logger().Debugf("playerModel: %+v", playerModel)

				if !playerModel.IsValid() {

					// NOTE: No Timeout used here
					ch <- errors.New("playerModel Validation failed")
					return
				}

				c.Logger().Debug("playerModel is valid passing it to the Player handler")
				handlers.Player(*playerModel, c) //TODO: handle errors
			}
		}
	}

}
