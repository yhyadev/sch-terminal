package websocket

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func WriteMessages(code string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		content, err := reader.ReadString('\n')

		if err != nil {
			panic(err.Error())
		}

		message, _ := json.Marshal(map[string]interface{}{
			"type":      "message",
			"content":   strings.TrimSpace(content),
			"room_code": code,
		})

		server.Write(message)

		fmt.Print("\033[1A")
	}
}
