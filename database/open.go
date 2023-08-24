package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func open() *gorm.DB {

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

func Init() {
	DB = open()
}

func GetDB() *gorm.DB {
	return DB
}
