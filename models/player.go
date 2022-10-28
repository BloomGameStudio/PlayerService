package models

import (
	uuid "github.com/satori/go.uuid"
)

type Player struct {
	UserID   uuid.UUID `gorm:"type:uuid;primarykey"`
	Location Coordinate
	Rotation Rotation
	State    State
}
