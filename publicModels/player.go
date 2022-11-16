package publicModels

import "github.com/BloomGameStudio/PlayerService/models"

// OPTIMIZE: Staticly link requestPlayer with Model.Player
type Player struct {
	// COMBAK: Add needed further fields from the Player struct model
	Name     string          `json:"name"`
	Position models.Position `json:"position"`
	Rotation models.Rotation `json:"rotation"`
	Scale    models.Scale    `json:"scale"`
}

func (rp Player) IsValid() bool {

	// Validates the requestPlayer
	// Additional validation and hooks for the reqeurequestPlayer validation can be added here
	// WARNING: Validation should be scoped to the requestPlayer

	if !rp.Position.IsValid() {
		return false
	}

	if !rp.Rotation.IsValid() {
		return false
	}

	if !rp.Scale.IsValid() {
		return false
	}

	return true
}
