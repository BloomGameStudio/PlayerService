package publicModels

type Transform struct {
	Position Position `json:"position" `
	Rotation Rotation `json:"rotation" `
	Scale    Scale    `json:"scale"`
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
