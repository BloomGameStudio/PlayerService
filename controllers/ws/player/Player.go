package player

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func Player(c echo.Context) error {

	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
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
