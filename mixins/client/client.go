package client

import (
	"github.com/gorilla/websocket"
)

// ConditionalWriter executes WriteFunc depending on the SendData flag.
func ConditionalWriter(socket *websocket.Conn, sendData bool, writeFunc func() error) error {
	if sendData {
		err := writeFunc()
		if err != nil {
			// Handle error
			return err
		}
		// Optionally handle success or return nil
		return nil
	} else {
		// Handle when sendData is false
		return nil
	}
}