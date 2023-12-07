package mixin

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

func (c *ClientMixin) ConditionalWriter(socket *websocket.Conn, writeFunc func() error) {
	if c.SendData {
		err := writeFunc()
		if err != nil {
			//Handle error
		}
	} else {
		//Handle when sendData is false
	}
}
