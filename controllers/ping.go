package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingResp struct {
	Ping string `json:"ping" form:"ping"`
}

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &PingResp{"pong"})
}
