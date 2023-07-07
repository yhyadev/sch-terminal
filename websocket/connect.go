package websocket

import (
	"golang.org/x/net/websocket"
)

func Connect() *websocket.Conn {
  conn, err := websocket.Dial("ws://sch.nawafdev.com/ws", "", "http://localhost")

  if err != nil {
    panic(err.Error())
  }

  return conn
}
