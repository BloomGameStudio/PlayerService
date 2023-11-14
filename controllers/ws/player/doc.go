// Package player provides a WebSocket controller for handling player-related operations in the playerService.
// It includes functions for upgrading the connection to a WebSocket, reading JSON data from the WebSocket,
// validating the received data, and writing JSON data to the WebSocket based on updates from the database.
// The package consists of the following files:
//   - `player.go`: Contains the `Player` function that is the entry point for handling WebSocket communication.
//   - `playerReader.go`: Contains the `playerReader` function that reads JSON data from the WebSocket and processes it.
//   - `playerWriter.go`: Contains the `playerWriter` function that writes JSON data to the WebSocket based on database updates.
//
// Please note that this package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws
// - github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers
// - github.com/BloomGameStudio/PlayerService/database
// - github.com/BloomGameStudio/PlayerService/handlers
// - github.com/BloomGameStudio/PlayerService/models
// - github.com/BloomGameStudio/PlayerService/publicModels
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
// - github.com/spf13/viper
// - gorm.io/gorm/clause
//
// For more information on each function, please refer to their individual documentation.
package player