package playerModel

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
)

func PlayerModel(c echo.Context) error {

	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	go playerModelWriter(c, ws, writerChan, timeoutCTX)
	go playerModelReader(c, ws, readerChan, timeoutCTX)

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
