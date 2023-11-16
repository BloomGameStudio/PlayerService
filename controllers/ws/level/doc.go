// Package level provides a WebSocket controller for handling the "level" event in the playerService.
//
// This event is used for level-related operations in the game. The `Level` function is the entry point for handling
// GET requests on the "ws/level" endpoint. It connects to the WebSocket and performs the following steps:
//   1. Upgrades the connection to a WebSocket.
//   2. Reads JSON data from the WebSocket and parses it into level-related models.
//   3. Validates the received data and calls the appropriate handlers.
//   4. Writes JSON data to the WebSocket based on updates from the database.
//   5. Sends periodic ping messages to the client for WebSocket health checks.
//
// The `level` package includes the following files:
//   - `level.go`: Contains the `Level` function that handles the WebSocket communication and dispatches tasks.
//   - `levelReader.go`: Contains the `levelReader` function that reads JSON data from the WebSocket and processes it.
//   - `levelWriter.go`: Contains the `levelWriter` function that writes JSON data to the WebSocket based on database updates.
//
// Please note that this package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
// - github.com/spf13/viper
//
// For more information on each function, please refer to their individual documentation.
package level