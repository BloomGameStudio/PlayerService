package models

type Transform struct {
	RotationID uint
	ScaleID    uint
	PositionID uint
	Position   Position `json:"position" `
	Rotation   Rotation `json:"rotation" `
	Scale      Scale    `json:"scale"`
	// publicModels.Transform
}

// func (s TransformIDS) IsValid() bool {

// 	return true
// }
