package ws

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	GetRate = func(c echo.Context) int {
		params := c.Request().URL.Query()
		rateStr := params.Get("rate")
		rate, err := strconv.Atoi(rateStr)
		if err != nil {
			rate = 1
		}
		return rate
	}
)
