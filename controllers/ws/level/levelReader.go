package level

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

func levelReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:

			// Initializer model to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqLevel := &models.Level{}

			err := ws.ReadJSON(reqLevel)

			if err != nil {
				errorHandlers.HandleReadError(c, ch, err, false)
			}

			if !reqLevel.IsValid() {
				// NOTE: no Chan Timeout used
				ch <- errors.New("reqLevel Validation failed")
				return
			}

			// Use dot annotation for promoted aka embedded fields.
			levelModel := &models.Level{}
			// TODO: Handle ID and production mode

			if viper.GetBool("DEBUG") {
				// Accept client provided ID and PlayerID in DEBUG mode
				levelModel.ID = reqLevel.ID
				levelModel.PlayerID = reqLevel.PlayerID

			}

			levelModel.LevelID = reqLevel.LevelID

			if !levelModel.IsValid() {
				// NOTE: no Chan Timeout used
				ch <- errors.New("levelModel Validation failed")
				return
			}

			c.Logger().Debug("rotationModel is valid passing it to the Poisition handler")
			handlers.Level(*levelModel, c)
		}
	}
}
