// The Cartesian coordinate and the Bearing system is used
// https://en.wikipedia.org/wiki/Cartesian_coordinate_system
// https://en.wikipedia.org/wiki/Bearing_(angle)

package models

type Orientation struct {

	// Represent orientation in 3D
	// e.g A VerticalBearing and a HoriontalBearing of 0
	// would mean that something is orientated north on the Horizontal plain
	// while also being parallel to the Horizontal plain

	// A VerticalBare of 090 would be facing north while being parallel with the Z axes which most people would
	// consider looking/facing up

	// Very simply put a circle is overlayed over a grid
	// Good real life representation is a  "Gyro/Sphere Compass"

	HorizontalBearing uint16 // The bearing on the Horizontal axes
	// On a 2D grid a bearing of 000 would align perfectly on the Y axes and means the Orientation is North
	// 180 would mean South
	// A bearing of 090 would align perfectly with the X axes and means the Orientation is East
	// NOTE: while a bearing of 090 definitely means the Orientation is East
	// Movement can not be derived from it as it is unkown

	VerticalBearing uint16 // The bearing on the Vertical axes
	// Exact same as HorizontalBearing just on the Vectical axes
	// bearing of 000 will have orientation facing north(On a 2D grid not a 3D) parallel to the Horizontal plane
	// most people would consider the following:
	// 000 == Facing forward
	// 090 = Facing upward
	// 180 = Facing backwards
	// and so on
}
