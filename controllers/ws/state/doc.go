// Package state provides functions for handling state-related operations in the playerService.
//
// The state package includes the following files:
//   - state.go: Contains the State function, which is the entry point for handling state WebSocket communication.
//   - stateReader.go: Contains the stateReader function, responsible for reading JSON data from the WebSocket, validating it, and invoking the corresponding handler.
//   - stateWriter.go: Contains the stateWriter function, responsible for retrieving state data from the database, sending it to the WebSocket, and handling error conditions.
//
// Dependencies:
// The state package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws
// - github.com/BloomGameStudio/PlayerService/models
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
// - github.com/spf13/viper
//
// For more information on each function, please refer to their individual documentation.
package state