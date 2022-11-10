package models

import "gorm.io/gorm"

type Vector3 struct {
	gorm.Model
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
