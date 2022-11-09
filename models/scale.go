package models

import "gorm.io/gorm"

type Scale struct {
	gorm.Model
	XYZ
}

func (s Scale) IsValid() bool {

	return true
}
