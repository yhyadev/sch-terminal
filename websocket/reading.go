package websocket

import (
	"fmt"

	"github.com/yhyadev/sch-terminal/message"
)

func ReadMessages() {
	go func() {
		for {
			bytes := make([]byte, 4096)

			n, err := server.Read(bytes)
			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("%s\n", message.Parse(n, bytes))
		}
	}()
}
