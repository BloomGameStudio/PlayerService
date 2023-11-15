package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	//Cannot embed publicModels.Model as Model already declared.
	//Explicitly define fields, this still allows them to be accessible
    ModelID    uint `json:"id"`
    MaterialID uint `json:"material_id"`
}

func (m Model) IsValid() bool {

	return true
}
