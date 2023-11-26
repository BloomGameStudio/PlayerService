package state

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func stateReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {
		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return
		default:
			c.Logger().Debug("Reading from the State WebSocket")

			// Initialize request states to bind into
			reqStates := []models.State{}

			err := ws.ReadJSON(&reqStates)

			if err != nil {
				switch err.(type) {
				case *json.UnmarshalTypeError:
					c.Logger().Error(err)
				default:
					errorHandlers.HandleReadError(c, ch, err)
					return
				}
			}

			for _, reqState := range reqStates {
				stateModel := &models.State{}

				if !reqState.IsValid() {
					ch <- errors.New("reqState validation failed")
					return
				}


				// Use dot annotation for promoted aka embedded fields.
				// TODO: Handle ID and production mode

				if viper.GetBool("DEBUG") {
					// Accept client provided ID in DEBUG mode
					stateModel.ID = reqState.ID
				}

				if reqState.ID <= 0 {
					ch <- errors.New("missing/invalid ID")
					return
				}

				stateModel.StateID = reqState.StateID
				stateModel.Value = reqState.Value

				// stateModel.Airborn = reqState.Airborn
				// stateModel.Grounded = reqState.Grounded
				// stateModel.Waterborn = reqState.Waterborn

				if !stateModel.IsValid() {
					ch <- errors.New("stateModel validation failed")
					return
				}


				// Pass each reqState individually to handlers.State
				err := handlers.State(reqState, c)
				if err != nil {
					ch <- err
					return
				}
			}
		}
	}
}