package main

import "fmt"

func main() {
  fmt.Println("SCH - For Terminals")
  fmt.Println()

  code := roomPicker()
  fmt.Print(code)
}

func roomPicker() string {
  var code string

  fmt.Print("Pick A Room (Room Code) : ")

  fmt.Scanln(&code)

  return code
}
