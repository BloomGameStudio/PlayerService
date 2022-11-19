package publicModels

import "gorm.io/gorm"

type Scale struct {
	gorm.Model
	Vector3
}

func (s Scale) IsValid() bool {

	return true
}
