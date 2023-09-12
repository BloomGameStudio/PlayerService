package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// For some reason that is beyond me the Gorm "Has one" relationship does not work
// Therefore we have to use the "Belongs To" relationship
type Player struct {
	// models.Player holds private fields for the player model

	gorm.Model

	UserID uuid.UUID `gorm:"type:uuid;uniqueIndex"`

	publicModels.Player

	Transform `json:"transform"`
	States    []State `json:"states"`
	publicModels.PModel
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
