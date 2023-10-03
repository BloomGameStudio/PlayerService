package publicModels

type Rotation struct {
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll
	Vector3
	W      float64 `json:"w"`
	Active bool    `json:"active"` // default to true

	// EulerAngles *Vector3 `json:"eulerAngles"`
	// Perhaps we want to have verbose EulerAngles that point to a Vector3
}

func (r Rotation) IsValid() bool {

	return true
}
