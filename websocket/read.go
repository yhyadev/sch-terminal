package websocket

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func Parse(n int, message []byte) string {
	var parsed Message

	err := json.Unmarshal(message[:n], &parsed)
	if err != nil {
		panic(err.Error())
	}

	return parsed.Content
}

func ReadMessages() {
	go func() {
		for {
			bytes := make([]byte, 4096)

			n, err := server.Read(bytes)
			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("%s\n", Parse(n, bytes))
		}
	}()
}
