// Package playerservice 
// This package is responsible for handling players, including
// creating, retrieving, updating, and deleting player information.
// The player service provides a RESTful API to interact with player data
// websocket connections are also supported.
//
// Subpackages of the application include:
// - config:      Configuration management and setup
// - controllers: Application logic and handler mapping
// - database:    Database integration and management
// - docs:        Generated documentation for the API and examples
// - handlers:    handlers for endpoints
// - helpers:     Utility functions that assist in various tasks
// - models:      Data structures for database entities
// - publicModels: Data structures for publicly exposed entities
//
// Usage
//
// The service can be started by running command ```go run .``` 
// this will setup the necessary routes and start listening
// for HTTP requests. Alternatively you can run ```docker-compose up```
//
// The API is then accessible through localhost on port 1323
package playerservice
