package publicModels

type States struct {
	IDS string `json:"ids"`
}

type Player struct {
	// publicModels.Player holds public fields for the player model.

	// COMBAK: Add needed further fields from the Player struct model
	Name      string `json:"name" gorm:"uniqueIndex"`
	Layer     string `json:"layer"`
	ENS       string `json:"ens"`
	Active    bool   `json:"active" gorm:"default:1"` // default to true only tested on SQLite might behave differently on other databases
	Transform `json:"transform" gorm:"-:all"`

	States `json:"states"`
}

func (p Player) IsValid() bool {

	// Validates the public Player
	// Additional validation and hooks for the public Player validation can be added here
	// WARNING: Validation should be scoped to the public Player

	return true
}
