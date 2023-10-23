package publicModels

type Player struct {
	// publicModels.Player holds public fields for the player model.

	// COMBAK: Add needed further fields from the Player struct model
	Name   string `json:"name" gorm:"uniqueIndex"`
	Layer  string `json:"layer"`
	ENS    string `json:"ens"`
	Active *bool   `json:"active,omitempty" gorm:"default:1"` // default to true only tested on SQLite might behave differently on other databases

	Level     Level `json:"level" gorm:"-:all"`
	Transform `json:"transform" gorm:"-:all"`
	States    []State `json:"states" gorm:"-:all"`
}

func (p Player) IsValid() bool {

	// Validates the public Player
	// Additional validation and hooks for the public Player validation can be added here
	// WARNING: Validation should be scoped to the public Player

	return true
}
