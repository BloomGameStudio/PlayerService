package state

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

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

	select {
	case w := <-writerChan:
		c.Logger().Debugf("Received writerChan error: %v", w)
		return nil
	case r := <-readerChan:
		c.Logger().Debugf("Received readerChan error: %v", r)
		return nil
	}
}
