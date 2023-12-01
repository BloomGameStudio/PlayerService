// Package playerModel provides a WebSocket controller for handling the player model in the playerService.
// This event is used for updating and retrieving player models in the game.

// The playerModel package includes the following files:
//   - playerModel.go: contains the main PlayerModel function that handles the model event.
//   - playerModelReader.go: contains the playerModelReader function that reads the model data from the WebSocket connection.
//   - playerModelWriter.go: contains the playerModelWriter function that writes the model data to the WebSocket connection.

// The PlayerModel function is responsible for upgrading the HTTP connection to a WebSocket connection and initializing the playerModelReader and playerModelWriter goroutines to handle the communication.

// The playerModelReader function reads the model data sent by the client from the WebSocket connection, validates it, and passes it to the playerModel handler for further processing.

// The playerModelWriter function periodically retrieves the latest models from the database and sends the model data to the client via the WebSocket connection. It also performs a ping check to ensure the WebSocket connection is still active.

// Note: The playerModel package relies on other packages and imports external dependencies such as gorilla/websocket, labstack/echo, BloomGameStudio/PlayerService/handlers, BloomGameStudio/PlayerService/models, and BloomGameStudio/PlayerService/database.
package playerModel