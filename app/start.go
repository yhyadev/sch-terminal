package app

import (
	"github.com/yhyadev/sch-terminal/websocket"
)

func Start(code string, nickname string) {
	server := websocket.Connect()
	defer server.Close()

	websocket.JoinRoom(code, nickname)

	websocket.ReadMessages()

	websocket.WriteMessages(code)
}
