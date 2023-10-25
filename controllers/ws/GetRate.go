package ws

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRate(c echo.Context) int {
	rateStr := c.QueryParam("rate")
	rate, err := strconv.Atoi(rateStr)
	if err != nil {
		rate = 1
	}
	return rate
}
