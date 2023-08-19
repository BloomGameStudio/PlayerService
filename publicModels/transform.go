package publicModels

type Transform struct {
	PositionID uint
	Position   Position `json:"position" `
	RotationID uint
	Rotation   Rotation `json:"rotation" `
	ScaleID    uint
	Scale      Scale `json:"scale"`
}

func (s Transform) IsValid() bool {

	if !s.Position.IsValid() {
		return false
	}

	if !s.Rotation.IsValid() {
		return false
	}

	if !s.Scale.IsValid() {
		return false
	}

	return true
}
