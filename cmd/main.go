package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "naration.json", "The file to read the json object from.")
	flag.Parse()

	fmt.Printf("Reading from file %s\n", *filename)
}
