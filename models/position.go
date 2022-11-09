package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	XYZ
	// Normalized   *Position `json:"normalized,omitempty"`
	// Magnitude    float64   `json:"magnitude"`
	// SqrMagnitude float64   `json:"sqrMagnitude"`
}

func (p Position) IsValid() bool {

	return true
}
