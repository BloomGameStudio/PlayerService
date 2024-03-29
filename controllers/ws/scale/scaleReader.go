package scale

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

func scaleReader(
	c echo.Context,
	ws *websocket.Conn,
	ch chan error,
	timeoutContext context.Context,
) {
	for {
		select {
		case <-timeoutContext.Done():
			return
		default:
			reqScale := &models.Scale{}

			err := ws.ReadJSON(reqScale)

			if err != nil {
				errorHandlers.HandleReadError(c, ch, err)
			}

			if !reqScale.IsValid() {
				ch <- errors.New("reqScale validation failed")
				return
			}

			scaleModel := &models.Scale{}

			// Accept client provided ID in DEBUG mode
			if viper.GetBool("DEBUG") {
				scaleModel.ID = reqScale.ID
			}

			scaleModel.Vector3 = reqScale.Vector3

			if reqScale.ID <= 0 {
				ch <- errors.New("missing/invalid ID")
				return
			}

			if !scaleModel.IsValid() {
				ch <- errors.New("scaleModel validation failed")
				return
			}

			handlers.Scale(*scaleModel, c)
		}
	}
}
