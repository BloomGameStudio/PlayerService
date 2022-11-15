package controllers

import (
	"github.com/BloomGameStudio/PlayerService/handlers"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func Player(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		func() {

			// Initializer request player to bind into
			reqPlayer := &RequestPlayer{}

			err := ws.ReadJSON(reqPlayer)

			if err != nil {
				c.Logger().Error(err)
			}

			handlers.Player(reqPlayer)

		}()

	}
}
