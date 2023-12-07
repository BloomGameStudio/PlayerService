package client

import (
	"github.com/gorilla/websocket"
)

type ClientMixin struct {
	SendData bool
}

// Create an instance of ClientMixin and accept a boolean as the parameter
func NewClientMixin(sendData bool) *ClientMixin {
	//Return the derefenced value of the ClientMixin boolean
	return &ClientMixin{SendData: sendData}
}

// Conditional execution of WriteFunc depending on SendData flag.
func (c *ClientMixin) ConditionalWriter(socket *websocket.Conn, WriteFunc func() error) {
	if c.SendData {
		err := WriteFunc()
		if err != nil {
			//Handle error
		}
	} else {
		//Handle when sendData is false
	}
}
