package ws

import (
	"strconv"

	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

type Model interface {
	GetPosition() models.Position
}

func RadiusFilter[m Model](a *[]m, c echo.Context) *[]m {
	radius, err_radius := strconv.ParseFloat(c.QueryParam("radius"), 32)
	anchorPointX, err_anchorPointX := strconv.ParseFloat(c.QueryParam("anchorPointX"), 32)
	anchorPointY, err_anchorPointY := strconv.ParseFloat(c.QueryParam("anchorPointY"), 32)

	if err_radius == nil && err_anchorPointX == nil && err_anchorPointY == nil {
		// valid parameters were provided

		// filters positions slice in-place according to radius

		// can use generics to have func accept multiple types like this:
		// func sumIntsOrFloats[V int | float64](m []V) V {
		//
		// }

		p := (*a)[:0]

		for i := range *a {
			if Distance(anchorPointX, anchorPointY, (*a)[i].GetPosition().X, (*a)[i].GetPosition().Y) < radius {
				p = append(p, (*a)[i])
			}
		}

		return &p

	}
	return a

}
