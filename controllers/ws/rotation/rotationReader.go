package rotation

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

func rotationReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the Rotation WebSocket")

			// Initializer request player to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqRotation := &models.Rotation{}

			err := ws.ReadJSON(reqRotation)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqRotation")
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

			c.Logger().Debugf("reqRotation from the WebSocket: %+v", reqRotation)

			c.Logger().Debug("Validating reqRotation")
			if !reqRotation.IsValid() {
				c.Logger().Debug("reqRotation is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("reqRotation Validation failed")
				return
			}

			c.Logger().Debug("reqRotation is valid")

			c.Logger().Debug("Initializing and populating rotation model!")
			// Use dot annotation for promoted aka embedded fields.
			rotationModel := &models.Rotation{}
			// TODO: Handle ID and production mode

			if viper.GetBool("DEBUG") {
				// Accept client provided ID in DEBUG mode
				rotationModel.ID = reqRotation.ID
			}

			rotationModel.Vector3 = reqRotation.Vector3
			rotationModel.W = reqRotation.W

			c.Logger().Debugf("rotationModel: %+v", rotationModel)

			c.Logger().Debug("Validating rotationModel")
			if !rotationModel.IsValid() {
				c.Logger().Debug("rotationModel is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("rotationModel Validation failed")
				return
			}

			c.Logger().Debug("rotationModel is valid passing it to the Poisition handler")
			handlers.Rotation(*rotationModel, c)
		}
	}
}
