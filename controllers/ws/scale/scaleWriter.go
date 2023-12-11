package scale

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/BloomGameStudio/PlayerService/mixins/client"
)

func scaleWriter(
	c echo.Context,
	ws *websocket.Conn,
	ch chan error,
	timeoutContext context.Context,
	sendData bool) {
	db := database.GetDB()
	lastUpdatedAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	lastPingCheck := time.Now()
	wsTimeout := time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

	for {
		select {
		case <-timeoutContext.Done():
			return
		default:
			scales := &[]models.Scale{}
			db.Where("updated_at > ?", lastUpdatedAt).Find(scales)
			lastUpdatedAt = time.Now()

			if len(*scales) > 0 {

				err := client.ConditionalWriter(ws, sendData, func() error {
					if sendData {
						return ws.WriteJSON(scales)
					}
					// Optionally handle case when sendData is false
					return nil
				})

				if err != nil {
					switch {
					case errors.Is(err, websocket.ErrCloseSent):
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
			} else if lastPingCheck.Add(time.Second).Before(time.Now()) {
				err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))

				if err != nil {
					switch {
					case errors.Is(err, websocket.ErrCloseSent):
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
			}

			time.Sleep(time.Millisecond)

			// Sleep for x second in DEBUG mode to not get fludded with data
			if viper.GetBool("DEBUG") {
				time.Sleep(time.Second / 20)
			}
		}
	}
}
