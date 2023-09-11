package controllers

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func Rotation(c echo.Context) error {

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

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	go rotationWriter(c, ws, writerChan, timeoutCTX)
	go rotationReader(c, ws, readerChan, timeoutCTX)

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

// Read
func rotationReader(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {

	// TODO: THIS IS VULNARABLE CLIENTS CAN CHANGE OBJECT IDS especially the nested ones!!!
	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!

	for {
		select {

		case <-timeoutCTX.Done():
			c.Logger().Debug("TimeoutCTX Done")
			return

		default:
			c.Logger().Debug("Reading from the Rotation WebSocket")

			// Initializer request player to bind into
			// NOTE: We are using a private model here TODO: Change to public model in production or handle this case
			reqRotation := &models.Rotation{}

			err := ws.ReadJSON(reqRotation)

			if err != nil {
				c.Logger().Debug("We get an error from Reading the JSON reqRotation")
				switch {

				case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
					select {

					case ch <- nil:
						c.Logger().Debug("Sent nil to Reader channel")
						return

					case <-time.After(time.Second * 10):
						c.Logger().Debug("Timed out sending nil to Reader channel")
						return
					}

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
				}
			}

			c.Logger().Debugf("reqRotation from the WebSocket: %+v", reqRotation)

			c.Logger().Debug("Validating reqRotation")
			if !reqRotation.IsValid() {
				c.Logger().Debug("reqRotation is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("reqRotation Validation failed")
				return
			}

			c.Logger().Debug("reqRotation is valid")

			c.Logger().Debug("Initializing and populating rotation model!")
			// Use dot annotation for promoted aka embedded fields.
			rotationModel := &models.Rotation{}
			// TODO: Handle ID and production mode

			if viper.GetBool("DEBUG") {
				// Accept client provided ID in DEBUG mode
				rotationModel.ID = reqRotation.ID
			}

			rotationModel.Vector3 = reqRotation.Vector3
			rotationModel.W = reqRotation.W

			c.Logger().Debugf("rotationModel: %+v", rotationModel)

			c.Logger().Debug("Validating rotationModel")
			if !rotationModel.IsValid() {
				c.Logger().Debug("rotationModel is NOT valid returning")
				// NOTE: no Chan Timeout used
				ch <- errors.New("rotationModel Validation failed")
				return
			}

			c.Logger().Debug("rotationModel is valid passing it to the Poisition handler")
			handlers.Rotation(*rotationModel, c)
		}
	}
}
