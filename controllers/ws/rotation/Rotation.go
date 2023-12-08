package rotation

import (
	"context"
	"fmt"
	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
	"strconv" // Import strconv for string to boolean conversion
)

// NOTE: We may need to adjust default configuration and values
// examples:
// https://github.com/gorilla/websocket/blob/master/examples/command/main.go

func Rotation(c echo.Context) error {
	// Extract the sendData value from the query parameters
	sendDataStr := c.QueryParam("sendData")

	// Convert the string to a boolean
	sendData, err := strconv.ParseBool(sendDataStr)
	if err != nil {
		return fmt.Errorf("failed to parse sendData: %v", err)
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

	go rotationWriter(c, ws, writerChan, timeoutCTX, sendData)
	go rotationReader(c, ws, readerChan, timeoutCTX)

	// QUESTION: Do we want to wait on both routines to error out?
	// Return nil if either the reader or the writer encounters a error
	// Do NOT return the error this will cause the error "the connection is hijacked"

	select {

	case w := <-writerChan:
		c.Logger().Debugf("Received writerChan error: %v", w)
		return nil

	case r := <-readerChan:
		c.Logger().Debugf("Received readerChan error: %v", r)
		return nil

	}
}
