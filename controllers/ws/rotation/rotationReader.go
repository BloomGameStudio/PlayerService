package rotation

import (
	"context"
	"errors"

	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
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
				errorHandlers.HandleReadError(c, ch, err)
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
			rotationModel.ID = reqRotation.ID

			if rotationModel.ID <= 0 {
				ch <- errors.New("missing/invalid ID")
				return
			}

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
