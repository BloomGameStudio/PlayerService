package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Health struct {
	gorm.Model
	publicModels.Attribute
}

func (h Health) IsValid() bool {

	return true
}
