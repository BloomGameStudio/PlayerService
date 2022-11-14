package models

type Coordinate struct {
	Location Vector3 
	region int //should this be a Region type instead, or an int mapping to such
}

func (c Coordinate) IsValid() bool {
	if c.region.IsValid(Location) {
		return true
	}
}
