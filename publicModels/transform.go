package publicModels

import "gorm.io/gorm"

type Transform struct {
	gorm.Model
}

func (s Transform) IsValid() bool {

	return true
}
