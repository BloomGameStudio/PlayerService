package position

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

func positionWriter(c echo.Context, socket *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// Open DB outside of the loop
	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()
	wsTimeout := time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("PositionWriter Timeout Context Done")
			return

		default:
			// TODO: The Entire Position Model is being sent. It may contain information that should not be sent!

			c.Logger().Debug("Getting positions from the database")

			positions := &[]models.Position{}

			db.Where("updated_at > ?", lastUpdateAt).Find(positions)
			lastUpdateAt = time.Now() // update last update time to now only included positions that have been updated

			positions = ws.RadiusFilter(positions, c)

			if len(*positions) > 0 {

				c.Logger().Debug("Pushing the positions to the WebSocket")
				err := socket.WriteJSON(positions)

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

				// Run Ping Check if there are no results to send and last ping check was older than 1 second ago
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

			// Update Interval NOTE: setting depending on the server and its performance either increase or decrease it.
			// rate query param passed in by client, set to 1 by default

			rate := ws.GetRate(c)
			time.Sleep(time.Millisecond * time.Duration(rate))

			if viper.GetBool("DEBUG") {
				// Sleep for x second in DEBUG mode to not get fludded with data
				time.Sleep(time.Second / 20)
			}
		}
	}

}
