package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transform struct {
	gorm.Model
	RotationID uint
	ScaleID    uint
	PositionID uint
	PlayerID   uuid.UUID
	Position   Position `json:"position" `
	Rotation   Rotation `json:"rotation" `
	Scale      Scale    `json:"scale"`
	// publicModels.Transform
}

// func (s TransformIDS) IsValid() bool {

// 	return true
// }
