package publicModels

type Scale struct {
	Vector3
}

func (s Scale) IsValid() bool {
	return true
}
