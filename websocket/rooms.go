package websocket

import (
	"encoding/json"
	"fmt"
)

func JoinRoom(code string, nickname string) {
	go func() {
		room, _ := json.Marshal(map[string]interface{}{
			"type":      "room_join",
			"room_code": code,
			"nickname":  nickname,
		})

		server.Write(room)
		fmt.Println()
	}()
}
