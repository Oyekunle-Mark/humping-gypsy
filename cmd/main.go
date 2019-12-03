package main

import (
	"encoding/json"
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

	d := json.NewDecoder(f)

	var story gypsy.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
