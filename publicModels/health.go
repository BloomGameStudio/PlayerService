package publicModels

type Health struct {
	Attribute
}

func (h Health) IsValid() bool {

	return true
}
