package ws

import (
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

type Model_R interface {
	GetRotation() models.Rotation
}

func RotationFilter[m Model_R](a *[]m, c echo.Context) *[]m {

	// valid parameters were provided

	// filters positions slice in-place according to radius

	// can use generics to have func accept multiple types like this:
	// func sumIntsOrFloats[V int | float64](m []V) V {
	//
	// }

	r := (*a)[:0]

	for i := range *a {
		if *(*a)[i].GetRotation().Active {
			r = append(r, (*a)[i])
		}
	}

	return &r

}
