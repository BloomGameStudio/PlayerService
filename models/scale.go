package models

type Scale struct {
	XYZ
}

func (s Scale) IsValid() bool {

	return true
}
