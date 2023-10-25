package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	PublicModel publicModels.Model
}

func (m Model) isValid() bool {
	
	return true
}