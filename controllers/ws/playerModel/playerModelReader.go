package playerModel

import (
	"context"
	"errors"

	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
)

func playerModelReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {
	for {
		select {
		case <-timeoutCTX.Done():
			errorHandlers.SendNilOrTimeout(c, ch)
			return

		default:
			reqModel := &models.PlayerModel{}
			err := ws.ReadJSON(reqModel)
			if err != nil {
				errorHandlers.HandleReadError(c, ch, err)
				return
			}

			if !reqModel.IsValid() {
				errorHandlers.SendErrOrTimeout(c, ch, errors.New("reqModel validation failed"))
				return
			}

			model := &models.PlayerModel{}
			if viper.GetBool("DEBUG") {
				model.ID = reqModel.ID
			}
			
			model.MaterialID = reqModel.MaterialID

			if !model.IsValid() {
				errorHandlers.SendErrOrTimeout(c, ch, errors.New("model validation failed"))
				return
			}

			handlers.PlayerModel(*model, c)
		}
	}
}
