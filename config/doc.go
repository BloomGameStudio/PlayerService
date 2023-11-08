// Package config
// This package is responsibile for initializing the configuration of the playerService
// It sets default configuration values and initializes the database.
//
// Initialization
//
// The configuration is initialized when running docker compose .
// The following tasks are performed:
// Calls ViperInit function to set default configuration variables
// Calls the database.init function to initialize  the database connection
// and perform migrations
//
// Setting Default Values
//
// The 'ViperInit' function sets default values for certain configuration
// variables using the viper package
// These values will be used if no corresponding values are found in the configuration
// file or env variables
// The following configuration variables have default values set:
// - DEBUG:		true
// - PORT: 		"1323"
// - WS_TIMEOUT_SECONDS: 10
//
// Package Dependencies
//
// The config package depends on the following packages:
// - "github.com/BloomGameStudio/PlayerService/database": For initializing the database connection
// - "github.com/BloomGameStudio/PlayerService/models": For performing database migrations.
// - "github.com/spf13/viper": For setting default configuration values.
//
package config