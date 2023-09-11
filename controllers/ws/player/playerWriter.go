package player

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
	"gorm.io/gorm/clause"
)

func playerWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("PlayerWriter Timeout Context Done")
			return

		default:

			// TODO: The Entire Player Model is being sent. It may contain information that should not be sent!
			c.Logger().Debug("Writing to the WebSocket")

			c.Logger().Debug("Getting all players from the database")
			// Get all active players from the database
			queryPlayer := &models.Player{}
			queryPlayer.Active = true

			players := &[]models.Player{}

			db.Preload(clause.Associations).Where("updated_at > ?", lastUpdateAt).Where(queryPlayer).Find(players)
			lastUpdateAt = time.Now() // update last update time to now only included players that have been updated

			if len(*players) > 0 {

				// TODO: Find/Filter the Changes that occured in the players and send them NOTE: The above filters for changes pretty well but we may want to filter for specific changes
				// PlayerChanges(players,players)

				c.Logger().Debug("Pushing the player to the WebSocket")
				err := ws.WriteJSON(players)

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
