// Package controllers contains all the controller functions used by the application.
package controllers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"
)

// PingResp defines a structure that holds the ping value response.
type PingResp struct {
	Ping string `json:"ping" form:"ping"`
}

// Ping handles GET requests on the "/ping" endpoint.
// It returns a JSON response with the message "pong".
func Ping(c echo.Context) error {
	go func() {
		for {
			c.Logger().Debugf("Num of Goroutines: %v", runtime.NumGoroutine())

			time.Sleep(time.Second * 2)
		}

	}()
	return c.JSON(http.StatusOK, &PingResp{"pong"})
}
