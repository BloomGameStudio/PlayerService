package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Health struct {
	gorm.Model
	UserID uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	publicModels.Attribute
}

func (h Health) IsValid() bool {

	return true
}
