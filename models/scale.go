package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Scale struct {
	gorm.Model
	publicModels.Scale
}

func (s Scale) IsValid() bool {
	return true
}
