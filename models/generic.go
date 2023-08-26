package models

import (
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"gorm.io/gorm"
)

type Vector3 struct {
	gorm.Model
	publicModels.Vector3
}
