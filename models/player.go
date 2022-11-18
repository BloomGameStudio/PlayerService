package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// For some reason that is beyond me the Gorm "Has one" relationship does not work
// Therefore we have to use the "Belongs To" relationship
type Player struct {
	gorm.Model
	UserID     uuid.UUID `gorm:"type:uuid;index"`
	Name       string    `json:"name"`
	Layer      string    `json:"layer"`
	PositionID uint
	Position   Position `json:"position" `
	RotationID int
	Rotation   Rotation `json:"rotation" `
	ScaleID    int
	Scale      Scale `json:"scale"`
	// State    State
}

func (p Player) IsValid() bool {

	// Validates the Player
	// Additional validation and hooks for the Player validation can be added here
	// WARNING: Validation should be scoped to the Player

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
