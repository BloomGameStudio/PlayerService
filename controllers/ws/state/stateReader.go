package state

import (
	"context"
	"errors"

	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func stateReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!
forloop:
	for {
		c.Logger().Debug("Reading from the WebSocket")

		// Initializer request player to bind into
		reqState := &publicModels.State{}
		err := ws.ReadJSON(reqState)

		if err != nil {
			c.Logger().Debug("We get an error from Reading the JSON reqState")
			switch {

			case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
				c.Logger().Debug("Websocket CloseNoStatusReceived")
				ch <- nil
				close(ch)
				break forloop

			default:
				c.Logger().Error(err)
				ch <- err
				close(ch)
				break forloop

			}
		}

		c.Logger().Debugf("reqState from the WebSocket: %+v", reqState)

		c.Logger().Debug("Validating reqState")
		if !reqState.IsValid() {
			c.Logger().Debug("reqState is NOT valid returning")
			ch <- errors.New("reqState Validation failed")
			close(ch)
			break
		}

		c.Logger().Debug("reqState is valid")

		c.Logger().Debug("Initializing and populating reqState model!")
		// Use dot annotation for promoted aka embedded fields.
		stateModel := &models.State{}

		// stateModel.Airborn = reqState.Airborn
		// stateModel.Grounded = reqState.Grounded
		// stateModel.Waterborn = reqState.Waterborn

		if viper.GetBool("DEBUG") {

		}

		if reqState.ID <= 0 {
			ch <- errors.New("missing/invalid ID")
			return
		}

		c.Logger().Debugf("stateModel: %+v", stateModel)

		c.Logger().Debug("Validating stateModel")
		if !stateModel.IsValid() {
			c.Logger().Debug("stateModel is NOT valid returning")
			ch <- errors.New("stateModel Validation failed")
			close(ch)
			break
		}

		c.Logger().Debug("playerModel is valid passing it to the Player handler")
		// handlers.State(*stateModel, c) TODO: Implement StateHandler
	}

}
