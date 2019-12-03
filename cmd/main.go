package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "naration.json", "The file to read the json object from.")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic("Cannot read file")
	}

	fmt.Println(f)
}
