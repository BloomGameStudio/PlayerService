package model

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

func Model(c echo.Context) error {
	// Upgrade the HTTP server connection to the WebSocket protocol
	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error("Failed to set websocket upgrade:", err)
		return err
	}
	defer ws.Close()

	// Channels to receive errors (or nil) from reader/writer goroutines
	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	// Start the reader/writer goroutines
	go modelWriter(c, ws, writerChan, timeoutCTX)
	go modelReader(c, ws, readerChan, timeoutCTX)

	select {

	case w := <-writerChan:
		c.Logger().Debugf("Recieved writerChan error: %v", w)
		return nil

	case r := <-readerChan:
		c.Logger().Debugf("Recieved readerChan error: %v", r)
		return nil

	}
}
