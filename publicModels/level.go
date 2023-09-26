package publicModels

import "gorm.io/gorm"

type Level struct {
	// publicModels.State holds public fields for the State model.

	gorm.Model      // NOTE: COMEBACK: Accepting ID from Client Should only be for debug mode
	LevelID    uint `json:"levelID"`
}

func (l Level) IsValid() bool {

	// Validates the public State
	// Additional validation and hooks for the public State validation can be added here
	// WARNING: Validation should be scoped to the public State

	return true
}
