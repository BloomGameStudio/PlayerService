package rotation

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func rotationWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// Open DB outside of the loop
	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("RotationWriter Timeout Context Done")
			return

		default:
			// TODO: The Entire Rotation Model is being sent. It may contain information that should not be sent!

			c.Logger().Debug("Getting rotations from the database")

			rotations := &[]models.Rotation{}

			db.Where("updated_at > ?", lastUpdateAt).Find(rotations)
			lastUpdateAt = time.Now() // update last update time to now only included rotations that have been updated

			if len(*rotations) > 0 {

				c.Logger().Debug("Pushing the rotations to the WebSocket")
				err := ws.WriteJSON(rotations)

				if err != nil {
					switch {

					case errors.Is(err, websocket.ErrCloseSent):

						select {

						case ch <- nil:
							c.Logger().Debug("Sent nil to Writer channel")
							return

						case <-time.After(time.Second * 10):
							c.Logger().Debug("Timed out sending nil to Writer channel")
							return
						}

					default:
						c.Logger().Error(err)
						select {
						case ch <- err:
							c.Logger().Debug("Sent error to Writer channel")
							return

						case <-time.After(time.Second * 10):
							c.Logger().Debug("Timed out sending error to Writer channel")
							return
						}
					}
				}

				// Run Ping Check if there are no results to send and last ping check was older than 1 second ago
			} else if lastPingCheck.Add(time.Second * 1).Before(time.Now()) {
				c.Logger().Debug("Running Ping Check")

				err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))

				if err != nil {
					switch {

					case errors.Is(err, websocket.ErrCloseSent):
						c.Logger().Debug("WEbsocket ErrCloseSent")

						select {

						case ch <- nil:
							c.Logger().Debug("Sent nil to Writer channel")
							return
						case <-time.After(time.Second * 10):
							c.Logger().Debug("Timed out sending nil to Writer channel")
							return
						}

					default:
						c.Logger().Error(err)

						select {

						case ch <- err:
							c.Logger().Debug("Sent error to Writer channel")
							return
						case <-time.After(time.Second * 10):
							c.Logger().Debug("Timed out sending error to Writer channel")
							return
						}
					}
				}
			}

			c.Logger().Debug("Finished writing to the WebSocket Sleeping now")

			// Update Interval NOTE: setting depending on the server and its performance either increase or decrease it.
			// rate query param passed in by client, set to 1 by default

			params := c.Request().URL.Query()
			rateStr := params.Get("rate")
			rate, err := strconv.Atoi(rateStr)
			if err != nil {
				rate = 1
			}

			time.Sleep(time.Millisecond * time.Duration(rate))

			if viper.GetBool("DEBUG") {
				// Sleep for x second in DEBUG mode to not get fludded with data
				time.Sleep(time.Second / 20)
			}
		}
	}

}
