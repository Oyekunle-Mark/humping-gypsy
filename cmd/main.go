package main

import (
	"flag"
	"fmt"
	"os"

	gypsy "github.com/Oyekunle-Mark/humping-gypsy"
)

func main() {
	filename := flag.String("file", "naration.json", "The file to read the json object from.")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := gypsy.JsonStory(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
