package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	publicModels.Position
	// Normalized   *Position `json:"normalized,omitempty"`
	// Magnitude    float64   `json:"magnitude"`
	// SqrMagnitude float64   `json:"sqrMagnitude"`
}

func (p Position) IsValid() bool {

	return true
}
