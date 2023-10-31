package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ModelData publicModels.Model `json:"model"`
}

func (m Model) IsValid() bool {

	return true
}
