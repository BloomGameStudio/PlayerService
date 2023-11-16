// Package handlers provides handler functions for different operations related to players.
//
// The handlers package handles incoming HTTP requests and performs CRUD operations on the database. It includes
// handlers for the following operations:
//
//   - Level: Handles the level operation and updates the database with the provided level information.
//   - Player: Handles the player operation and updates the database with the provided player information.
//   - Position: Handles the position operation and updates the database with the provided position information.
//   - Rotation: Handles the rotation operation and updates the database with the provided rotation information.
//   - Scale: Handles the scale operation and updates the database with the provided scale information.
//   - playerModel: Handles the playerModel operation and updates the database with the provided playerModel information.   
//
// Files
//
// The package consists of the following files:
//
//   - level.go: Contains the Level function that handles the level operation.
//   - player.go: Contains the Player function that handles the player operation.
//   - position.go: Contains the Position function that handles the position operation.
//   - rotation.go: Contains the Rotation function that handles the rotation operation.
//   - scale.go: Contains the Scale function that handles the scale operation.
//	 - playerModel.go: Contains the playerModel function that handles the model operation.
// Dependencies
//
// The handlers package relies on the following packages:
//
//   - github.com/BloomGameStudio/PlayerService/controllers/ws: provides the WebSocket controller.
//   - github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers: handles WebSocket error handling.
//   - github.com/BloomGameStudio/PlayerService/database: provides database connection.
//   - github.com/BloomGameStudio/PlayerService/handlers: handles CRUD operations.
//   - github.com/BloomGameStudio/PlayerService/models: provides data models used in the player service.
//   - github.com/BloomGameStudio/PlayerService/publicModels: provides public data models used in the player service.
//
// Additionally, the handlers package relies on the following external packages:
//
//   - github.com/gorilla/websocket: handles the WebSocket connection and message handling.
//   - github.com/labstack/echo/v4: provides the web framework for the echo server.
//   - github.com/spf13/viper: handles configuration management.
//   - gorm.io/gorm/clause: provides additional query clauses for GORM.
//
// For more information on each function, please refer to their individual documentation.
package handlers