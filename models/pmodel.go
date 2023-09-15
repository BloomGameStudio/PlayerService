package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type PModel struct {
	gorm.Model
	publicModels.PModel
}

func (p PModel) isValid() bool {
	
	return true
}