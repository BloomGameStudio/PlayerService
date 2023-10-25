package state

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func stateWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: Retrivement of data needs to be defined
	// In Memory storage of the states has been agreed on
	// Open DB outside of the loop
	db := database.GetDB()
forloop:
	for {

		c.Logger().Debug("Writing to the WebSocket")
		c.Logger().Debug("Getting all States from the database")
		// Get all the states from the database
		states := &models.State{} // COMBAK: Data structure TBD
		db.Preload(clause.Associations).Find(states)

		// Find/Filter the Changes that occured in the states and send them

		c.Logger().Debug("Pushing the states to the WebSocket")
		err := ws.WriteJSON(states)
		if err != nil {

			switch {

			case errors.Is(err, websocket.ErrCloseSent):
				c.Logger().Debug("WEbsocket ErrCloseSent")
				ch <- nil
				close(ch)
				break forloop

			default:
				c.Logger().Error(err)
				ch <- err
				close(ch)
				break forloop
			}
		}
		c.Logger().Debug("Finished writing to the WebSocket Sleeping now")

		// Update Interval NOTE: setting depending on the server and its performance either increase or decrease it.
		time.Sleep(time.Millisecond * 1)

		if viper.GetBool("DEBUG") {
			// Sleep for 1 second in DEBUG mode to not get fludded with data
			time.Sleep(time.Second * 1)
		}
	}
}
