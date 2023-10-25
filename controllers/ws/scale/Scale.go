package scale

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

func Scale(c echo.Context) error {
	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutContext, timeoutContextCancel := context.WithCancel(context.Background())
	defer timeoutContextCancel()

	go scaleWriter(c, ws, writerChan, timeoutContext)
	go scaleReader(c, ws, readerChan, timeoutContext)

	select {
	case w := <-writerChan:
		c.Logger().Debugf("Recieved writerChan error: %v", w)
		return nil
	case r := <-readerChan:
		c.Logger().Debugf("Recieved readerChan error: %v", r)
		return nil
	}
}
