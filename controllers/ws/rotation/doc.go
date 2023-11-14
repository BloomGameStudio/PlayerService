// Package rotation provides functions for handling rotation-related operations in the playerService.
// It includes the `Rotation` function that is the entry point for handling WebSocket communication related to rotation.
// The package consists of the following files:
//   - `rotation.go`: Contains the `Rotation` function that upgrades the connection to a WebSocket, initializes reader and writer goroutines, and handles the overall rotation logic.
//   - `rotationReader.go`: Contains the `rotationReader` function that reads JSON data from the WebSocket, validates it, and invokes the corresponding handler.
//   - `rotationWriter.go`: Contains the `rotationWriter` function that retrieves rotation data from the database, sends it to the WebSocket, and handles error conditions.
//
// Please note that this package relies on the following packages:
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
package rotation