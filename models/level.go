package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	PlayerID uint
	publicModels.Level
}

func (l Level) IsValid() bool {

	// Validates the State
	// Additional validation and hooks for the State validation can be added here
	// WARNING: Validation should be scoped to the State

	return true
}
