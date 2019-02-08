package main

import (
	"fmt"
	"glob/tegj"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if arg != nil {
		resp := tegj.GenerateResponse(strings.Join(arg, " "))

		fmt.Println(resp)
	}
}
