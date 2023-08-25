package publicModels

type Rotation struct {
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll
	Vector3
	W float64 `json:"w"`

	// EulerAngles *Vector3 `json:"eulerAngles"`
	// Perhaps we want to have verbose EulerAngles that point to a Vector3
}

func (r Rotation) IsValid() bool {

	return true
}
