package models

import (
	"gorm.io/gorm"
	"github.com/BloomGameStudio/PlayerService/publicModels"
)

type PlayerModel struct {
	gorm.Model
    publicModels.PlayerModel
}

func (m PlayerModel) IsValid() bool {

	return true
}
