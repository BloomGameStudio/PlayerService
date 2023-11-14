package state

import (
	"context"
	"time"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func stateWriter(c echo.Context, socket *websocket.Conn, ch chan error, timeoutCTX context.Context) {
	// Open DB outside of the loop
	db := database.GetDB()
	//Use some very old date for the first update
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	lastPingCheck := time.Now()

	for {
		select {
		case <-timeoutCTX.Done():
			c.Logger().Debug("StateWriter Timeout Context Done")
			return
		default:
			c.Logger().Debug("Getting states from the database")

			states := &[]models.State{}

			db.Where("updated_at > ?", lastUpdateAt).Find(states)
			lastUpdateAt = time.Now()

			if len(*states) > 0 {
				c.Logger().Debug("Pushing the states to the WebSocket")
				err := socket.WriteJSON(states)

				if err != nil {
					errorHandlers.HandleWriteError(c, ch, err)
					return
				}
			} else if lastPingCheck.Add(time.Second * 1).Before(time.Now()) {
				err := socket.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))
				if err != nil {
					errorHandlers.HandleReadError(c, ch, err)
					return
				}
			}

			// Update Interval
			rate := ws.GetRate(c)
			time.Sleep(time.Millisecond * time.Duration(rate))

			if viper.GetBool("DEBUG") {
				time.Sleep(time.Second / 20)
			}
		}
	}
}
