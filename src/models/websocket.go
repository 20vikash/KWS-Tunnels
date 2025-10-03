package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSConn struct {
	Conn  *websocket.Conn
	Write *sync.Mutex
}
