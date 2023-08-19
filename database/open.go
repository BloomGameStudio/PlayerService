package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open() *gorm.DB {

	// memoryDB := "file::memory:?cache=shared"
	fileDB := "database.db"

	db, err := gorm.Open(sqlite.Open(fileDB), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true})
	if err != nil {
		panic("failed to connect database")
	}

	return db

}
