// Package player provides RESTful HTTP controllers for handling player-related endpoints in the playerService.
// This package includes controllers for retrieving player information, creating new players, updating existing players, and deleting players.

// CreatePlayer.go:
// The CreatePlayer function handles the creation of a new player in the service. It expects a request body containing either a publicModel or a model.Player object. 
// The function binds the request body to a reqPlayer variable, validates it, and then initializes and populates a playerModel.
// The playerModel is saved to the database, and the created player is returned in the HTTP response.

// DeletePlayer.go:
// The DeletePlayer function handles the deletion of a player based on a specified identifier. It retrieves the player from the database based on the identifier (UUID, ID, or name).
// If the player is not found, a 404 error is returned. If an error occurs during the database operation, a 500 error is returned.
// If the player is successfully deleted, a 200 status with the message "Player deleted successfully" is returned.

// GetPlayer.go:
// The GetPlayer function retrieves a list of players from the database based on the "active" query parameter.
// The function reads the "active" query parameter from the URL and parses it as a boolean value.
// It builds a query based on the "active" filter and retrieves a list of players from the database.
// The list of players is then returned in the HTTP response as a JSON array.

// UpdatePlayer.go:
// UpdatePlayer function handles the updating of player information. It takes the player identifier from the request path parameter.
// The function retrieves the player from the database based on the identifier (UUID, ID, or name).
// If the player is found, the JSON request body is parsed into an updatedData variable.
// The function then updates the specific fields in the queryPlayer object with the fields from updatedData.
// NOTE: If a field is NOT included it was be updated to a null / nil value.
// The updated player is saved to the database, and the updated player object is returned in the HTTP response as JSON.
// If the player is not found in the database, a 404 error is returned. If there is an error during the database operation, a 500 error is returned.


package player