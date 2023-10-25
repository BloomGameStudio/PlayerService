package player

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

func playerWriter(c echo.Context, socket *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()
	wsTimeout := time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

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
			active := true
			queryPlayer.Active = &active


			players := &[]models.Player{}

			db.Preload(clause.Associations).Where("updated_at > ?", lastUpdateAt).Where(queryPlayer).Find(players)
			lastUpdateAt = time.Now() // update last update time to now only included players that have been updated

			players = ws.RadiusFilter(players, c)

			if len(*players) > 0 {

				// TODO: Find/Filter the Changes that occured in the players and send them NOTE: The above filters for changes pretty well but we may want to filter for specific changes
				// PlayerChanges(players,players)

				c.Logger().Debug("Pushing the player to the WebSocket")
				err := socket.WriteJSON(players)

				if err != nil {
					errorHandlers.HandleWriteError(c, ch, err)
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
