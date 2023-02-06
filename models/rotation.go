package models

import "gorm.io/gorm"

type Rotation struct {
	gorm.Model
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll
	Vector3
	w float32

	// EulerAngles *Vector3 `json:"eulerAngles"`
	// Perhaps we want to have verbose EulerAngles that point to a Vector3
}

func (r Rotation) IsValid() bool {

	return true
}
