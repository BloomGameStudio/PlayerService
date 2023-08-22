package publicModels

import "gorm.io/gorm"

type Position struct {
	gorm.Model // TODO: Remove DB fields from public Model there might already be a branch for this
	Vector3
	// Normalized   *Position `json:"normalized,omitempty"`
	// Magnitude    float64   `json:"magnitude"`
	// SqrMagnitude float64   `json:"sqrMagnitude"`
}

func (p Position) IsValid() bool {

	return true
}
