package position

import (
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Distance(x1, y1, x2, y2 float64) float64 {
	distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return distance
}

func positionWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// Open DB outside of the loop
	db := database.GetDB()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("PositionWriter Timeout Context Done")
			return

		default:
			// TODO: The Entire Position Model is being sent. It may contain information that should not be sent!

			c.Logger().Debug("Getting positions from the database")

			positions := &[]models.Position{}

			params := c.Request().URL.Query()
			radiusStr := params.Get("radius")
			radius, err := strconv.ParseFloat(radiusStr, 32)
			if err != nil {
				// error handling, set radius to 1.0 as default value
				radius = 1.0
			}

			var startingPoint_X float64 = 0
			var startingPoint_Y float64 = 0

			filter := func(pos *[]models.Position) *[]models.Position {
				var out []models.Position
				for i := 0; i < len(*pos); i++ {
					if Distance(startingPoint_X, startingPoint_Y, (*pos)[i].X, (*pos)[i].Y) >= radius {
						out = append(out, (*pos)[i])
					}
				}
				return &out
			}
			positions = filter(positions)

			db.Where("updated_at > ?", lastUpdateAt).Find(positions)
			lastUpdateAt = time.Now() // update last update time to now only included positions that have been updated

			// filteredPositions = filter(positions)

			if len(*positions) > 0 {

				c.Logger().Debug("Pushing the positions to the WebSocket")
				err := ws.WriteJSON(positions)

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
			time.Sleep(time.Millisecond * 1)

			if viper.GetBool("DEBUG") {
				// Sleep for x second in DEBUG mode to not get fludded with data
				time.Sleep(time.Second / 20)
			}
		}
	}

}
