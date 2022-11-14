package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Vector3
	// Normalized   *Position `json:"normalized,omitempty"`
	// Magnitude    float64   `json:"magnitude"`
	// SqrMagnitude float64   `json:"sqrMagnitude"`
}

func (r Region) IsValid(v Vector3) bool {

	return true
}
