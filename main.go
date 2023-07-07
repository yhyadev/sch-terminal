package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/yhyadev/sch-terminal/websocket"
)

func main() {
	fmt.Println("SCH - For Terminals")
	fmt.Println()

	code := roomPicker()
	nickname := nicknamePicker()

	startApp(code, nickname)
}

func roomPicker() string {
	var code string

	fmt.Print("Pick A Room (Room Code) : ")

	fmt.Scanln(&code)

	return code
}

func nicknamePicker() string {
	var nickname string

	fmt.Print("Pick A Nickname : ")

	fmt.Scanln(&nickname)

	return nickname
}

func startApp(code string, nickname string) {
	server := websocket.Connect()
	defer server.Close()

	go func() {
		room, _ := json.Marshal(map[string]interface{}{
			"type":      "room_join",
			"room_code": code,
			"nickname":  nickname,
		})

		server.Write(room)
		fmt.Println()
	}()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			message := make([]byte, 4096)

			n, err := server.Read(message)
			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("%s\n", parseMessae(n, message))
		}
	}()

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
func parseMessae(n int, message []byte) string {
	type Message struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}

	var parsed Message

	err := json.Unmarshal(message[:n], &parsed)
	if err != nil {
		panic(err.Error())
	}

	return parsed.Content
}
