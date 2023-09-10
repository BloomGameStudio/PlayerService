package ping

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PingResp defines a structure that holds the ping value response.
type PingResp struct {
	Ping string `json:"ping" form:"ping"`
}

// Ping handles GET requests on the "/ping" endpoint.
// It returns a JSON response with the message "pong".
func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &PingResp{"pong"})
}
