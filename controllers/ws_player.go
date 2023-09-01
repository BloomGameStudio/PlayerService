package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func Player(c echo.Context) error {

	// QUESTION: Is this needed?
	// Only changes will be sent the only exception to this is the opening/first request where the full state will be sent
	// Partial player data can be received or full
	// TODO: Partial Reads
	// TODO: Partial Writes

	// TODO: Finalize how IDs are expected and handled

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	go playerWriter(c, ws, writerChan, timeoutCTX)
	go playerReader(c, ws, readerChan, timeoutCTX)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return nil if either the reader or the writer encounters a error
	// Do NOT return the error this will cause the error "the connection is hijacked"

	select {

	case w := <-writerChan:
		c.Logger().Debugf("Recieved writerChan error: %v", w)
		return nil

	case r := <-readerChan:
		c.Logger().Debugf("Recieved readerChan error: %v", r)
		return nil

	}
}

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
						// close(ch)
						// c.Logger().Debug("Returning Now From Go Routine")
						// return

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

						// ch <- err
						// // close(ch)
						// c.Logger().Debug("Returning Now From Go Routine")
						// return
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
						// ch <- nil
						// // close(ch)
						// c.Logger().Debug("Returning Now From Go Routine")
						// return

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
						// ch <- err
						// // close(ch)
						// c.Logger().Debug("Returning Now From Go Routine")
						// return
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

// Read
func playerReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {

		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the WebSocket")

			// Initializer request player to bind into
			reqPlayer := &publicModels.Player{}
			err := ws.ReadJSON(reqPlayer)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqPlayer")
				switch {

				case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
					c.Logger().Debug("Websocket CloseNoStatusReceived")
					select {
					case ch <- nil:
						c.Logger().Debug("Sent nil to Reader channel")
						return

					case <-time.After(time.Second * 10):
						c.Logger().Debug("Timed out sending nil to Reader channel")
						return
					}
					// ch <- nil
					// // close(ch)
					// c.Logger().Debug("Returning Now From Reader Go Routine")
					// return

				default:
					c.Logger().Error(err)
					select {
					case ch <- err:
						c.Logger().Debug("Sent error to Reader channel")
						return
					case <-time.After(time.Second * 10):
						c.Logger().Debug("Timed out sending error to Reader channel")
						return
					}
					// ch <- err
					// // close(ch)
					// c.Logger().Debug("Returning Now From Reader Go Routine")
					// return
				}
			}

			c.Logger().Debugf("reqPlayer from the WebSocket: %+v", reqPlayer)

			c.Logger().Debug("Validating reqPlayer")
			if !reqPlayer.IsValid() {
				c.Logger().Debug("reqPlayer is NOT valid returning")
				ch <- errors.New("reqPlayer Validation failed")
				// close(ch)
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return
			}

			c.Logger().Debug("reqPlayer is valid")

			c.Logger().Debug("Initializing and populating player model!")
			// Use dot annotation for promoted aka embedded fields.
			playerModel := &models.Player{}
			// TODO: Handle UserID and production mode
			playerModel.Position.Position = reqPlayer.Position
			playerModel.Rotation.Rotation = reqPlayer.Rotation
			playerModel.Scale.Scale = reqPlayer.Scale
			playerModel.States = reqPlayer.States
			playerModel.Active = reqPlayer.Active

			if viper.GetBool("DEBUG") {
				// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
				playerModel.Name = reqPlayer.Name
			}

			c.Logger().Debugf("playerModel: %+v", playerModel)

			c.Logger().Debug("Validating playerModel")
			if !playerModel.IsValid() {
				c.Logger().Debug("playerModel is NOT valid returning")
				// NOTE: No Timeout used here
				ch <- errors.New("playerModel Validation failed")
				// close(ch)
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return
			}

			c.Logger().Debug("playerModel is valid passing it to the Player handler")
			handlers.Player(*playerModel, c) //TODO: handle errors
		}
	}

}
