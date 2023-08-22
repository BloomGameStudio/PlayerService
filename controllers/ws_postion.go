package controllers

import (
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func Position(c echo.Context) error {

	// QUESTION: Is this needed?
	// Only changes will be sent the only exception to this is the opening/first request where the full state will be sent
	// Partial player data can be received or full
	// TODO: Partial Reads
	// TODO: Partial Writes

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	go positionWriter(c, ws, writerChan)
	go positionReader(c, ws, readerChan)

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

// Write
func positionWriter(c echo.Context, ws *websocket.Conn, ch chan error) {

	// Open DB outside of the loop
	db := database.Open()
	lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC) // Use some ver old date for first update to get all players in the initial push
	lastPingCheck := time.Now()

	for {
		// TODO: The Entire Position Model is being sent. It may contain information that should not be sent!

		c.Logger().Debug("Getting positions from the database")

		positions := &[]models.Position{}

		db.Where("updated_at > ?", lastUpdateAt).Find(positions)
		lastUpdateAt = time.Now() // update last update time to now only included positions that have been updated

		if len(*positions) > 0 {

			c.Logger().Debug("Pushing the positions to the WebSocket")
			err := ws.WriteJSON(positions)

			if err != nil {
				switch {

				case errors.Is(err, websocket.ErrCloseSent):
					c.Logger().Debug("WEbsocket ErrCloseSent")
					ch <- nil
					return

				default:
					c.Logger().Error(err)
					ch <- err
					return
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
					ch <- nil
					return

				default:
					c.Logger().Error(err)
					ch <- err
					return
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

// Read
func positionReader(c echo.Context, ws *websocket.Conn, ch chan error) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		c.Logger().Debug("Reading from the WebSocket")

		// Initializer request player to bind into
		reqPlayer := &publicModels.Player{}
		err := ws.ReadJSON(reqPlayer)

		if err != nil {
			c.Logger().Debug("We get an error from Reading the JSON reqPlayer")
			switch {

			case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
				c.Logger().Debug("Websocket CloseNoStatusReceived")
				ch <- nil
				// close(ch)
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return

			default:
				c.Logger().Error(err)
				ch <- err
				// close(ch)
				c.Logger().Debug("Returning Now From Reader Go Routine")
				return
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
		playerModel.Position = reqPlayer.Position
		playerModel.Rotation = reqPlayer.Rotation
		playerModel.Scale = reqPlayer.Scale

		if viper.GetBool("DEBUG") {
			// Add the Player.Name in DEBUG mode that it can be used as ID in the Player handle to avoid the Userservice dependency
			playerModel.Name = reqPlayer.Name
		}

		c.Logger().Debugf("playerModel: %+v", playerModel)

		c.Logger().Debug("Validating playerModel")
		if !playerModel.IsValid() {
			c.Logger().Debug("playerModel is NOT valid returning")
			ch <- errors.New("playerModel Validation failed")
			// close(ch)
			c.Logger().Debug("Returning Now From Reader Go Routine")
			return
		}

		c.Logger().Debug("playerModel is valid passing it to the Player handler")
		handlers.Player(*playerModel, c) //TODO: handle errors
	}

}
