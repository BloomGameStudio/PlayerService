package errorHandlers

import (
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var wsTimeout time.Duration = time.Second * time.Duration(viper.GetInt("WS_TIMEOUT_SECONDS"))

func HandleWriteError(c echo.Context, ch chan error, err error) {

	c.Logger().Debug("We get an error from Writing the JSON")

	switch {

	case errors.Is(err, websocket.ErrCloseSent):
		HandleErrCloseSent(c, ch, err)
		return

	default:
		HandleUnknownError(c, ch, err)
		return

	}

}

func HandleReadError(c echo.Context, ch chan error, err error) {

	c.Logger().Debug("We get an error from Reading the JSON")

	switch {

	case websocket.IsCloseError(err, websocket.CloseNoStatusReceived):
		HandleCloseNoStatusReceived(c, ch)
		return

	default:
		HandleUnknownError(c, ch, err)
		return

	}

}

func HandleCloseNoStatusReceived(c echo.Context, ch chan error) {

	c.Logger().Debug("Websocket CloseNoStatusReceived")

	select {

	case ch <- nil:
		c.Logger().Debug("Sent nil to channel")
		return

	case <-time.After(wsTimeout):
		c.Logger().Debug("Timed out sending nil to channel")
		return

	}

}

func HandleErrCloseSent(c echo.Context, ch chan error, err error) {

	c.Logger().Debug("Websocket ErrCloseSent")

	select {

	case ch <- nil:
		c.Logger().Debug("Sent nil to Writer channel")
		return

	case <-time.After(wsTimeout):
		c.Logger().Debug("Timed out sending nil to Writer channel")
		return

	}
}

func HandleUnknownError(c echo.Context, ch chan error, err error) {

	c.Logger().Error(err)

	select {

	case ch <- err:
		c.Logger().Debug("Sent error to channel")
		return

	case <-time.After(wsTimeout):
		c.Logger().Debug("Timed out sending error to channel")
		return

	}

}
