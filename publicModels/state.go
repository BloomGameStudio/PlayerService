package publicModels

type State struct {
	// publicModels.State holds public fields for the State model.
	StateID uint    `json:"stateID"`
	Value   float64 `json:"value"`
	// Grounded, Airborn, Waterborn bool
}

// type States []State

// type States struct {
// 	States []State `json:"states"`
// }

// func (p State) IsValid() bool {

// 	// Validates the public State
// 	// Additional validation and hooks for the public State validation can be added here
// 	// WARNING: Validation should be scoped to the public State

// 	return true
// }
