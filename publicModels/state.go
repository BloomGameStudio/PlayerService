package publicModels

import "gorm.io/gorm"

type State struct {
	// publicModels.State holds public fields for the State model.

	gorm.Model         // NOTE: COMEBACK: Accepting ID from Client Should only be for debug mode
	StateID    uint    `json:"stateID"`
	Value      float64 `json:"value"`
	// Grounded, Airborn, Waterborn bool
}

func (p State) IsValid() bool {

	// Validates the public State
	// Additional validation and hooks for the public State validation can be added here
	// WARNING: Validation should be scoped to the public State

	return true
}
