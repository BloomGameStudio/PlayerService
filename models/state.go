package models

type State struct {
	Grounded, Airborn, Waterborn bool
}

func (p State) IsValid() bool {

	// Validates the State
	// Additional validation and hooks for the State validation can be added here
	// WARNING: Validation should be scoped to the State

	return true
}
