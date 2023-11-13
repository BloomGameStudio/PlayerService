// Package helpers provides utility functions for common operations used in the player service.
//
// The package consists of the following files:
//
//   - helpers.go: Contains helper functions for various tasks.
//
// Below are the helper functions available in
// the `helpers` package:
//
//   - GetPlayerModelFromJWT: Retrieves the player model from the JWT token in the provided echo context. This function
//                            returns the player model and an error.
//   - PrettyStruct: Beautifies a struct by marshaling it into an indented JSON string. This function takes a struct
//                   as input and returns the indented JSON string and an error.
//   - PrettyStructNoError: Beautifies a struct by marshaling it into an indented JSON string. This function takes a 
//                          struct as input and returns the indented JSON string. If an error occurs, it returns an 
//                          empty string.
//   - PrettyString: Beautifies a JSON string by indenting it using spaces. This function takes a string as input and
//                   returns the indented JSON string and an error.
//
// Dependencies
//
// The helpers package relies on the following packages:
//
//   - github.com/BloomGameStudio/PlayerService/database: provides database operations.
//   - github.com/BloomGameStudio/PlayerService/models: provides data models used in the player service.
//   - github.com/golang-jwt/jwt: handles JWT authentication and authorization.
//   - github.com/labstack/echo/v4: provides the web framework for the echo server.
//   - github.com/satori/go.uuid: generates and manipulates UUIDs.
//
// For more information on each function, please refer to their individual documentation.
package helpers