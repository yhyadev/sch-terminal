package app

import "fmt"

func RoomPicker() string {
	var code string

	fmt.Print("Pick A Room (Room Code) : ")

	fmt.Scanln(&code)

	return code
}

func NicknamePicker() string {
	var nickname string

	fmt.Print("Pick A Nickname : ")

	fmt.Scanln(&nickname)

	return nickname
}
