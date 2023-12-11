package playerModel

import (
	"context"

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
	"strconv"
)

func PlayerModel(c echo.Context) error {

	// Extract the sendData value from the query parameters
	sendDataStr := c.QueryParam("sendData")

	// Convert the string to a boolean
	sendData, err := strconv.ParseBool(sendDataStr)
	if err != nil {
		return err
	}

	ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	writerChan := make(chan error)
	readerChan := make(chan error)

	timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
	defer timeoutCTXCancel()

	go playerModelWriter(c, ws, writerChan, timeoutCTX, sendData)
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
