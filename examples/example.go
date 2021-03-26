package main

import (
	"fmt"

	"github.com/martinlindhe/inputbox"
)

func main() {
	got, ok := inputbox.InputBox("Dialog title", "Type a number", "abc")
	if ok {
		fmt.Println("you entered:", got)
	} else {
		fmt.Println("No value entered")
	}
}
