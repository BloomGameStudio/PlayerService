package ws

import (
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
)

func InactiveUpdateObjects(a *[]models.Player, c echo.Context) *[]models.Player {

	p := (*a)[:0]

	for i := range *a {
		p = append(p, (*a)[i])
		if !(p)[i].Player.Active {
			(p)[i].Transform.Rotation.Active = false
		}
	}

	return &p
}
