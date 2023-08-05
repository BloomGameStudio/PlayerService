// Package controllers contains all the controller functions used by the application.
package controllers

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// Hello handles GET requests on the "ws/hello" endpoint.
// It connects to the websocket and sends a message to the client.
func Hello(c echo.Context) error {
	// A Hello World type test websocket controller
	// reads a message from the websocket and prints it out to the console
	// writes Hello Client to the websocket

	// Upgrade the connection to a WebSocket
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
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		// Print the received message
		fmt.Printf("%s\n", msg)
	}
}
