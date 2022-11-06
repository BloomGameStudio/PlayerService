package models

import (
	uuid "github.com/satori/go.uuid"
)

type Player struct {
	UserID uuid.UUID `gorm:"type:uuid;primarykey"`
	// TODO: @Lapras What exactly does Name represent? Is it not always "GameObject"? and therefore could be infered? Or completly emited?
	Name string `json:"name"`
	// TODO: @Lapras Do we need a layer? Is the player not always on the PlayableCharacters layer or at least on the same consistante layer
	Layer    string   `json:"layer"`
	Position Position `json:"position"`
	Rotation Rotation `json:"rotation"`
	Scale    Scale    `json:"scale"`
	// State    State
}
