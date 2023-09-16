package version

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
)

type VersionResp struct {
	Version string `json:"version" form:"version"`
}

func Version(c echo.Context) error {
	return c.JSON(http.StatusOK, &VersionResp{runtime.Version()[2:]})
}

func Commit(c echo.Context) error {
	return c.JSON(http.StatusOK, &VersionResp{runtime.Version()[2:]})
}
