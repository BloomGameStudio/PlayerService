package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ModelID    uint
	MaterialID uint
}

func (m Model) IsValid() bool {

	return true
}
