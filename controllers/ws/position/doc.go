// Package position provides a WebSocket controller for handling the players position in the playerService.
// This event is used for updating and retrieving player positions in the game.

// The position package includes the following files:
//   - Position.go: contains the main Position function that handles the position event.
//   - positionReader.go: contains the positionReader function that reads the position data from the WebSocket connection.
//   - positionWriter.go: contains the positionWriter function that writes the position data to the WebSocket connection.

// The Position function is responsible for upgrading the HTTP connection to a WebSocket connection and initializing the positionReader and positionWriter goroutines to handle the communication.

// The positionReader function reads the position data sent by the client from the WebSocket connection, validates it, and passes it to the Position handler for further processing.

// The positionWriter function periodically retrieves the latest positions from the database, applies optional filtering based on radius, and sends the position data to the client via the WebSocket connection. It also performs a ping check to ensure the WebSocket connection is still active.

// Note: The position package relies on other packages and imports external dependencies such as gorilla/websocket and labstack/echo.
package position