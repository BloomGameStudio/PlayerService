package state

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func State(c echo.Context) error {

	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	go stateWriter(c, ws, writerChan, timeoutCTX)
	go stateReader(c, ws, readerChan, timeoutCTX)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return the error if either the reader or the writer encounters a error
	select {
	case w := <-writerChan:
		c.Logger().Debugf("Received writerChan error: %v", w)
		return nil
	case r := <-readerChan:
		c.Logger().Debugf("Received readerChan error: %v", r)
		return nil
	}
}
