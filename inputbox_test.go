package inputbox

import (
	"fmt"
	"testing"
)

func TestInputBox(t *testing.T) {
	got, ok := InputBox("Dialog title", "Type a number", "abc")
	if ok {
		fmt.Println("you entered:", got)
	} else {
		fmt.Println("No value entered")
	}
}
