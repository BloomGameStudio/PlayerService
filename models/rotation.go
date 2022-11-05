package models

type Rotation struct {
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll // TODO: Does this order reflect the client and is correct?
	XYZ

	// EulerAngles *XYZ `json:"eulerAngles"`
	// Perhaps we want to have verbose EulerAngles that point to XYZ
}
