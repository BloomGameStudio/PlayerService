package state

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func stateReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the States WebSocket")

			// Initializer request player to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqState := &models.State{}

			err := ws.ReadJSON(reqState)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqState")
				switch {

				case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
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

			c.Logger().Debugf("reqState from the WebSocket: %+v", reqState)

			c.Logger().Debug("Validating reqState")
			if !reqState.IsValid() {
				c.Logger().Debug("reqState is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("reqState Validation failed")
				return
			}

			c.Logger().Debug("reqState is valid")

			c.Logger().Debug("Initializing and populating state model!")
			// Use dot annotation for promoted aka embedded fields.
			stateModel := &models.State{}
			// TODO: Handle ID and production mode

			if viper.GetBool("DEBUG") {
				// Accept client provided ID in DEBUG mode
				stateModel.ID = reqState.ID
			}

			stateModel.StateID = reqState.StateID
			stateModel.Value = reqState.Value
			stateModel.Grounded = reqState.Grounded
			stateModel.Airborn = reqState.Airborn
			stateModel.Waterborn = reqState.Waterborn
			c.Logger().Debugf("stateModel: %+v", stateModel)

			c.Logger().Debug("Validating stateModel")
			if !stateModel.IsValid() {
				c.Logger().Debug("stateModel is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("stateModel Validation failed")
				return
			}

			c.Logger().Debug("stateModel is valid passing it to the state handler")
			handlers.State(*stateModel, c)
		}
	}
}
