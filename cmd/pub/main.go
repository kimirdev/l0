package main

import (
	"fmt"
	"l0/pub"
	"os"
)

var usage = `use valid\invalid argumnet
'valid' - publish valid model
'invalid' - publish invalid model

'./pub [valid/invalid]'`

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}

	switch os.Args[1] {
	case "valid":
		pub.PublishValid()
	case "invalid":
		pub.PublishInvalid()
	default:
		fmt.Println(usage)
	}
}
