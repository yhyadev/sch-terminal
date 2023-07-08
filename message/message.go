package message

import "encoding/json"

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
