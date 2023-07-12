package publicModels

import "gorm.io/gorm"

type Vector3 struct {
	gorm.Model
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Attribute struct {
	Ceiling   int `json:"ceiling"`
	Current   int `json:"current"`
	Collected int `json:"collected"`
}
