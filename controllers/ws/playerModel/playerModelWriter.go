package model

import (
	"context"
	"errors"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func playerModelWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context) {
    db := database.GetDB()
    lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
    lastPingCheck := time.Now()

    for {
        select {
        case <-timeoutCTX.Done():
            c.Logger().Debug("ModelWriter Timeout Context Done")
            return

        default:
            c.Logger().Debug("Getting models from the database")

            modelsList := &[]models.PlayerModel{}
            result := db.Where("updated_at > ?", lastUpdateAt).Find(modelsList) // Use result to capture DB operation details
            lastUpdateAt = time.Now()

            if result.Error != nil {
                c.Logger().Errorf("Database query error: %v", result.Error)
                ch <- result.Error
                return
            }

            c.Logger().Debugf("Number of models fetched: %d", len(*modelsList))
            if len(*modelsList) > 0 {
                for _, model := range *modelsList {
                    c.Logger().Debugf("Fetched model: %+v", model)
                }

                c.Logger().Debug("Pushing the models to the WebSocket")
                err := ws.WriteJSON(modelsList)
                if err != nil {
                    switch {
                    case errors.Is(err, websocket.ErrCloseSent):
                        select {
                        case ch <- nil:
                            c.Logger().Debug("Sent nil to Writer channel")
                            return
                        case <-time.After(time.Second * 10):
                            c.Logger().Debug("Timed out sending nil to Writer channel")
                            return
                        }

                    default:
                        c.Logger().Error(err)
                        select {
                        case ch <- err:
                            c.Logger().Debug("Sent error to Writer channel")
                            return
                        case <-time.After(time.Second * 10):
                            c.Logger().Debug("Timed out sending error to Writer channel")
                            return
                        }
                    }
                }
            }

            if lastPingCheck.Add(time.Second * 1).Before(time.Now()) {
                c.Logger().Debug("Running Ping Check")

                err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))
                if err != nil {
                    switch {
                    case errors.Is(err, websocket.ErrCloseSent):
                        c.Logger().Debug("WebSocket ErrCloseSent")

                        select {
                        case ch <- nil:
                            c.Logger().Debug("Sent nil to Writer channel")
                            return
                        case <-time.After(time.Second * 10):
                            c.Logger().Debug("Timed out sending nil to Writer channel")
                            return
                        }

                    default:
                        c.Logger().Error(err)

                        select {
                        case ch <- err:
                            c.Logger().Debug("Sent error to Writer channel")
                            return
                        case <-time.After(time.Second * 10):
                            c.Logger().Debug("Timed out sending error to Writer channel")
                            return
                        }
                    }
                }

                lastPingCheck = time.Now()
            }

            c.Logger().Debug("Finished accessing the database. Sleeping")
            time.Sleep(time.Millisecond * 1)

            if viper.GetBool("DEBUG") {
                time.Sleep(time.Second / 20)
            }
        }
    }
}
