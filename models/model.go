package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Data publicModels.Model
}

func (m Model) IsValid() bool {

	return true
}
