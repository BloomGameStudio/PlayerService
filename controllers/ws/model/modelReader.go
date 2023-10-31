package model

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

func modelReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {
	for {
		select {
		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the Model WebSocket")

			reqModel := &models.Model{}
			err := ws.ReadJSON(reqModel)
			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqModel")
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

			c.Logger().Debugf("reqModel from the WebSocket: %+v", reqModel)

			if !reqModel.IsValid() {
				c.Logger().Debug("reqModel is NOT valid returning")
				ch <- errors.New("reqModel Validation failed")
				return
			}

			c.Logger().Debug("reqModel is valid")

			model := &models.Model{}
			if viper.GetBool("DEBUG") {
				model.ID = reqModel.ID
			}
			model.ModelData = reqModel.ModelData

			c.Logger().Debugf("model: %+v", model)

			if !model.IsValid() {
				c.Logger().Debug("model is NOT valid returning")
				ch <- errors.New("model Validation failed")
				return
			}

			c.Logger().Debug("model is valid passing it to the Model handler")
			handlers.ModelHandler(*model, c)
		}
	}
}
