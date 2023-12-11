package playerModel

import (
	"context"
	"time"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/BloomGameStudio/PlayerService/controllers/ws/errorHandlers"
    "github.com/BloomGameStudio/PlayerService/mixins/client"
)

func playerModelWriter(c echo.Context, ws *websocket.Conn, ch chan error, timeoutCTX context.Context, sendData bool) {
    db := database.GetDB()
    lastUpdateAt := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
    lastPingCheck := time.Now()

    for {
        select {
        case <-timeoutCTX.Done():
            errorHandlers.SendNilOrTimeout(c, ch)
            return

        default:
            modelsList := &[]models.PlayerModel{}
            result := db.Where("updated_at > ?", lastUpdateAt).Find(modelsList)
            lastUpdateAt = time.Now()

            if result.Error != nil {
                errorHandlers.SendErrOrTimeout(c, ch, result.Error)
                return
            }

            if len(*modelsList) > 0 {

                c.Logger().Debug("Pushing the PlayerModels to the WebSocket")

                err := client.ConditionalWriter(ws, sendData, func() error {
                    if sendData {
                        return ws.WriteJSON(modelsList)
                    }
                    // Optionally handle case when sendData is false
                    return nil
                })

                if err != nil {
                    errorHandlers.HandleWriteError(c, ch, err)
                    return
                }
            }

            if lastPingCheck.Add(time.Second * 1).Before(time.Now()) {
                err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*2))
                if err != nil {
                    errorHandlers.HandleWriteError(c, ch, err)
                    return
                }
                lastPingCheck = time.Now()
            }

            sleepDuration := time.Millisecond * 1
            if viper.GetBool("DEBUG") {
                sleepDuration = time.Second / 20
            }
            time.Sleep(sleepDuration)
        }
    }
}
