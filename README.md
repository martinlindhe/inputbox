# About

[![GoDoc](https://godoc.org/github.com/martinlindhe/inputbox?status.svg)](https://godoc.org/github.com/martinlindhe/inputbox)

InputBox is a simple cross-platform library for displaying desktop dialog boxes in your go application


## Example

```go
package main

import "github.com/martinlindhe/inputbox"

func main() {
	got, ok := inputbox.InputBox("Dialog title", "Type a number", "abc")
	if ok {
		fmt.Println("you entered:", got)
	} else {
		fmt.Println("No value entered")
	}
}
```

### Windows

Uses a powershell script

![Windows](windows.png)

### macOS

Uses a osascript

![macOS](macos.png)


### Linux

Uses zenity (gtk)

![Linux](linux.png)


### More

If you like this, check out https://github.com/martinlindhe/notify for cross-platform desktop notifications.


### License

Under [MIT](LICENSE)
