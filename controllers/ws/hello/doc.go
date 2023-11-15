// Package hello provides a WebSocket controller for handling the "hello" event in the playerService.
//
// This event is used for greeting the client when they connect via WebSocket. The `Hello` function handles GET requests
// on the "ws/hello" endpoint. It connects to the websocket and sends a message to the client.
//
// The `Hello` function performs the following steps:
//   1. Upgrades the connection to a WebSocket.
//   2. Sends a "Hello, Client!" message to the client using the WebSocket.
//   3. Reads the response message from the client and logs it to the console.
//
// Please note that this package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
package hello