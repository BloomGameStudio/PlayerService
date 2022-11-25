package publicModels

type Player struct {
	// publicModels.Player holds public fields for the player model.

	// COMBAK: Add needed further fields from the Player struct model
	Name       string `json:"name" gorm:"uniqueIndex"`
	Layer      string `json:"layer"`
	PositionID uint
	Position   Position `json:"position" `
	RotationID uint
	Rotation   Rotation `json:"rotation" `
	ScaleID    uint
	Scale      Scale  `json:"scale"`
	ENS        string `json:"ens"`
}

func (p Player) IsValid() bool {

	// Validates the public Player
	// Additional validation and hooks for the public Player validation can be added here
	// WARNING: Validation should be scoped to the public Player

	if !p.Position.IsValid() {
		return false
	}

	if !p.Rotation.IsValid() {
		return false
	}

	if !p.Scale.IsValid() {
		return false
	}

	return true
}
