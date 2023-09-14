package scale

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
				wsTimeout := time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

				switch {
				case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
					select {
					case ch <- nil:
						return
					case <-time.After(wsTimeout):
						return
					}
				default:
					select {
					case ch <- err:
						return
					case <-time.After(wsTimeout):
						return
					}
				}
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

			if !scaleModel.IsValid() {
				ch <- errors.New("scaleModel validation failed")
				return
			}

			handlers.Scale(*scaleModel, c)
		}
	}
}
