package publicModels

type State struct {
	// publicModels.State holds public fields for the State model.

	Grounded, Airborn, Waterborn bool
}

func (p State) IsValid() bool {

	// Validates the public State
	// Additional validation and hooks for the public State validation can be added here
	// WARNING: Validation should be scoped to the public State

	return true
}
