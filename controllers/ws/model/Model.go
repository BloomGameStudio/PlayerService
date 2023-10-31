package model

import (
	"context"
	"net/http"
	"time"

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

	// Context to handle the timeout. Adjust accordingly to fit your needs.
	timeoutDuration := 60 * time.Minute
	timeoutCTX, timeoutCTXCancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer timeoutCTXCancel()

	// Start the reader/writer goroutines
	go modelWriter(c, ws, writerChan, timeoutCTX)
	go modelReader(c, ws, readerChan, timeoutCTX)

	select {
	case err := <-writerChan:
		if err != nil {
			c.Logger().Error("Model WebSocket Writer Error:", err)
		}
		return err
	case err := <-readerChan:
		if err != nil {
			c.Logger().Error("Model WebSocket Reader Error:", err)
		}
		return err
	case <-timeoutCTX.Done():
		c.Logger().Info("Model WebSocket connection closed due to timeout")
		return c.NoContent(http.StatusOK)
	}
}
