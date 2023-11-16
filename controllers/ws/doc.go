// Package ws contains the WebSocket controllers for handling real-time communication with clients in the playerService.
// It consists of sub-packages for different WebSocket events and functions.
//
// Sub-packages:
// - errorHandlers: Contains error handling functions for WebSocket communication.
// - hello: Contains WebSocket controllers for the "hello" event.
// - level: Contains WebSocket controllers for the "level" event.
// - player: Contains WebSocket controllers for the "player" event.
// - position: Contains WebSocket controllers for the "position" event.
// - rotation: Contains WebSocket controllers for the "rotation" event.
// - scale: Contains WebSocket controllers for the "scale" event.
// - state: Contains WebSocket controllers for the "state" event.
//
// Files:
// - Distance.go: Contains the Distance function, which calculates the distance between two points in a Cartesian coordinate system.
// - GetRate.go: Contains the GetRate function, which retrieves the "rate" query parameter from an HTTP request and converts it to an integer.
// - RadiusFilter.go: Contains the RadiusFilter function, which filters a slice of models based on a given radius and anchor point.
// - Upgrader.go: Contains the Upgrader variable, which is a websocket.Upgrader instance used to upgrade HTTP connections to WebSocket connections.
//
// Dependencies:
// The ws package relies on the following packages:
// - github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers
// - github.com/BloomGameStudio/PlayerService/controllers/ws/hello
// - github.com/BloomGameStudio/PlayerService/controllers/ws/level
// - github.com/BloomGameStudio/PlayerService/controllers/ws/player
// - github.com/BloomGameStudio/PlayerService/controllers/ws/position
// - github.com/BloomGameStudio/PlayerService/controllers/ws/rotation
// - github.com/BloomGameStudio/PlayerService/controllers/ws/scale
// - github.com/BloomGameStudio/PlayerService/controllers/ws/state
// - github.com/BloomGameStudio/PlayerService/models
// And the following external packages:
// - github.com/gorilla/websocket
// - github.com/labstack/echo/v4
// - github.com/spf13/viper
//
// For more information on each dependency, please refer to their individual documentation.
// For more information on each sub-package and file, please refer to their individual documentation.
//
package ws