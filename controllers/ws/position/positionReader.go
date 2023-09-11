package position

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

func positionReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the Position WebSocket")

			// Initializer request player to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqPosition := &models.Position{}

			err := ws.ReadJSON(reqPosition)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqPosition")
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

			c.Logger().Debugf("reqPosition from the WebSocket: %+v", reqPosition)

			c.Logger().Debug("Validating reqPosition")
			if !reqPosition.IsValid() {
				c.Logger().Debug("reqPosition is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("reqPosition Validation failed")
				return
			}

			c.Logger().Debug("reqPosition is valid")

			c.Logger().Debug("Initializing and populating position model!")
			// Use dot annotation for promoted aka embedded fields.
			positionModel := &models.Position{}
			// TODO: Handle ID and production mode

			if viper.GetBool("DEBUG") {
				// Accept client provided ID in DEBUG mode
				positionModel.ID = reqPosition.ID
			}

			positionModel.Vector3 = reqPosition.Vector3

			c.Logger().Debugf("positionModel: %+v", positionModel)

			c.Logger().Debug("Validating positionModel")
			if !positionModel.IsValid() {
				c.Logger().Debug("positionModel is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("positionModel Validation failed")
				return
			}

			c.Logger().Debug("positionModel is valid passing it to the Poisition handler")
			handlers.Position(*positionModel, c)
		}
	}
}
