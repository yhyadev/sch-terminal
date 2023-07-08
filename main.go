package main

import (
	"fmt"

	"github.com/yhyadev/sch-terminal/app"
)

func main() {
	fmt.Println("SCH - For Terminals")
	fmt.Println()

	code := app.RoomPicker()
	nickname := app.NicknamePicker()

	app.Start(code, nickname)
}
