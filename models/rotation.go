package models

type Rotation struct {
	// Euler Rotation Annotation
	// Yaw,Pitch,Roll TODO: Does this order reflect the client and is correct?
	X, Y, Z float64
}
