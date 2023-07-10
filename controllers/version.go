package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type VersionResp struct {
	Version string `json:"version" form:"version"`
}

func Version(c echo.Context) error {
	return c.JSON(http.StatusOK, &VersionResp{"0.0.1"})
}

func Commit(c echo.Context) error {
	return c.JSON(http.StatusOK, &VersionResp{"0.0.1"})
}
