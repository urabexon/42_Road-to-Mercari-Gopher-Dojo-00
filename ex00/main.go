package main

import (
	"fmt"
	"os"
)

func main() {
	// args check
	if len(os.Args) != 2 {
		fmt.Println("error: invalid argument")
		return
	}
	fmt.Println("Hello World!")
}
