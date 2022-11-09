package models

import "gorm.io/gorm"

type Rotation struct {
	gorm.Model
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll // TODO: @Lapras Does this order reflect the client and is correct?
	XYZ

	// EulerAngles *XYZ `json:"eulerAngles"`
	// Perhaps we want to have verbose EulerAngles that point to XYZ
}

func (r Rotation) IsValid() bool {

	return true
}
