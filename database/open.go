package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() *gorm.DB {

	// memoryDB := "file::memory:?cache=shared"
	fileDB := "database/database.db"

	db, err := gorm.Open(sqlite.Open(fileDB), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// Logger:                 logger.Default.LogMode(logger.Info), // Show Gorm SQL Queries
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db

}
