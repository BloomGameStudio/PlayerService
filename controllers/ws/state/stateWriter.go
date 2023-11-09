package state

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func stateWriter(c echo.Context, socket *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	lastPingCheck := time.Now()
	wsTimeout := time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

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
					switch {
					case errors.Is(err, websocket.ErrCloseSent):
						select {
						case ch <- nil:
							c.Logger().Debug("Sent nil to Writer channel")
							return
						case <-time.After(wsTimeout):
							c.Logger().Debug("Timed out sending nil to Writer channel")
							return
						}
					default:
						c.Logger().Error(err)
						select {
						case ch <- err:
							c.Logger().Debug("Sent error to Writer channel")
							return
						case <-time.After(wsTimeout):
							c.Logger().Debug("Timed out sending error to Writer channel")
							return
						}
					}
				}

			} else if lastPingCheck.Add(time.Second * 1).Before(time.Now()) {
				c.Logger().Debug("Running Ping Check")

				err := socket.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))

				if err != nil {
					switch {
					case errors.Is(err, websocket.ErrCloseSent):
						c.Logger().Debug("WEbsocket ErrCloseSent")
						select {
						case ch <- nil:
							c.Logger().Debug("Sent nil to Writer channel")
							return
						case <-time.After(wsTimeout):
							c.Logger().Debug("Timed out sending nil to Writer channel")
							return
						}
					default:
						c.Logger().Error(err)
						select {
						case ch <- err:
							c.Logger().Debug("Sent error to Writer channel")
							return
						case <-time.After(wsTimeout):
							c.Logger().Debug("Timed out sending error to Writer channel")
							return
						}
					}
				}
			}

			c.Logger().Debug("Finished writing to the WebSocket Sleeping now")

			// Update Interval
			rate := ws.GetRate(c)
			time.Sleep(time.Millisecond * time.Duration(rate))

			if viper.GetBool("DEBUG") {
				// Sleep for x second in DEBUG mode to not get flooded with data
				time.Sleep(time.Second / 20)
			}
		}
	}
}