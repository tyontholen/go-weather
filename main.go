package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: weather <city>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	city := os.Args[1]
	fmt.Println("You asked for the weather in:", city)

	// TODO: make api call
}
