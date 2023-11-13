// Package scale provides functions for handling scale-related operations in the playerService.
//
// The scale package includes the following files:
//   - scale.go: Contains the Scale function, which is the entry point for handling scale WebSocket communication.
//   - scaleReader.go: Contains the scaleReader function, responsible for reading JSON data from the WebSocket, validating it, and invoking the corresponding handler.
//   - scaleWriter.go: Contains the scaleWriter function, responsible for retrieving scale data from the database, sending it to the WebSocket, and handling error conditions.
//
// Dependencies:
// The scale package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws
// - github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers
// - github.com/BloomGameStudio/PlayerService/database
// - github.com/BloomGameStudio/PlayerService/handlers
// - github.com/BloomGameStudio/PlayerService/models
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
// - github.com/spf13/viper
//
// For more information on each function, please refer to their individual documentation.
package scale