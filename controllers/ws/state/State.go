package state

import (

	"github.com/BloomGameStudio/PlayerService/controllers/ws"
	"github.com/labstack/echo/v4"
	"context"

)
  //NOTE: We may need to adjust default configuration and values
  //examples:
  //https:github.com/gorilla/websocket/blob/master/examples/command/main.go

  func State(c echo.Context) error {
    // Create a context with cancellation
    timeoutCTX, timeoutCTXCancel := context.WithCancel(context.Background())
    defer timeoutCTXCancel()

    ws, err := ws.Upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()

    writerChan := make(chan error)
    readerChan := make(chan error)

    go stateWriter(c, ws, writerChan, timeoutCTX)
    go stateReader(c, ws, readerChan, timeoutCTX)

    // Wait for either the reader or the writer to encounter an error or for the context to be canceled
    select {
    case r := <-readerChan:
        c.Logger().Debugf("Received readerChan error: %v", r)
        return r
    case w := <-writerChan:
        c.Logger().Debugf("Received writerChan error: %v", w)
        return w
    case <-timeoutCTX.Done():
        // Context canceled, clean up gracefully
        return nil
    }
}
