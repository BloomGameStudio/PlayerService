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
	Scale      Scale `json:"scale"`
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
