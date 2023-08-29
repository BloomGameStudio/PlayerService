package publicModels

type Vector3 struct {
	// gorm.Model         // HACK: TODO: Remove this and derive fields This is for debug mode only people can provide IDs
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
