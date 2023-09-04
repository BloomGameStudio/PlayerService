package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	PlayerID uint
	publicModels.State

	// Grounded, Airborn, Waterborn bool
}

// type States []State

// type States struct {
// 	States []State `json:"states"`
// }

// func (p State) IsValid() bool {

// 	// Validates the State
// 	// Additional validation and hooks for the State validation can be added here
// 	// WARNING: Validation should be scoped to the State

// 	return true
// }
