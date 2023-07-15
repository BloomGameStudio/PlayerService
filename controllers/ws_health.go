package controllers

import (
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
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

func Health(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	go healthWriter(c, ws, writerChan)
	go healthReader(c, ws, readerChan)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return the error if either the reader or the writer encounters a error
	for {
		select {
		case r := <-readerChan:
			c.Logger().Debugf("Recieved readerChan error: %v", r)
			return r
		case w := <-writerChan:
			c.Logger().Debugf("Recieved writerChan error: %v", w)
			return w
		}
	}
}

// Write
func healthWriter(c echo.Context, ws *websocket.Conn, ch chan error) {

	// TODO: Retrivement of data needs to be defined
	// In Memory storage of the states has been agreed on
	// Open DB outside of the loop
	db := database.Open() // COMBAK: Configure the database
forloop:
	for {

		c.Logger().Debug("Writing to the WebSocket")
		c.Logger().Debug("Getting Health from the database")
		// Get all the states from the database
		health := &models.Health{} // COMBAK: Data structure TBD
		db.Preload(clause.Associations).Find(health)

		// Find/Filter the Changes that occured in the states and send them

		c.Logger().Debug("Pushing the states to the WebSocket")
		err := ws.WriteJSON(health)
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

// Read
func healthReader(c echo.Context, ws *websocket.Conn, ch chan error) {

	// TODO: NO VALIDATION OF INPUT DATA IS PERFORMED!!!
forloop:
	for {
		c.Logger().Debug("Reading from the WebSocket")

		// Initializer request player to bind into
		reqHealth := &publicModels.Health{}
		err := ws.ReadJSON(reqHealth)

		if err != nil {
			c.Logger().Debug("We get an error from Reading the JSON reqState")
			switch {

			case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
				c.Logger().Debug("Websocket CloseNoStatusReceived")
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

		c.Logger().Debugf("reqHealth from the WebSocket: %+v", reqHealth)

		c.Logger().Debug("Validating reqHealth")
		if !reqHealth.IsValid() {
			c.Logger().Debug("reqHealth is NOT valid returning")
			ch <- errors.New("reqHealth Validation failed")
			close(ch)
			break
		}

		c.Logger().Debug("reqHealth is valid")

		c.Logger().Debug("Initializing and populating reqHealth model!")
		// Use dot annotation for promoted aka embedded fields.
		healthModel := &models.Health{}

		healthModel.Attribute.Ceiling = reqHealth.Attribute.Ceiling
		healthModel.Attribute.Current = reqHealth.Attribute.Current
		healthModel.Attribute.Collected = reqHealth.Attribute.Collected

		if viper.GetBool("DEBUG") {

		}

		c.Logger().Debugf("healthModel: %+v", healthModel)

		c.Logger().Debug("Validating stateModel")
		if !healthModel.IsValid() {
			c.Logger().Debug("healthModel is NOT valid returning")
			ch <- errors.New("healthModel Validation failed")
			close(ch)
			break
		}

		c.Logger().Debug("healthModel is valid passing it to the Health handler")
		// handlers.Health(*healthModel, c) TODO: Implement healthHandler
	}

}
