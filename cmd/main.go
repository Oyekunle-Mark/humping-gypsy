package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	gypsy "github.com/Oyekunle-Mark/humping-gypsy"
)

func main() {
	filename := flag.String("file", "naration.json", "The file to read the json object from.")
	port := flag.Int("port", 5000, "The port to start the server on")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := gypsy.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := gypsy.NewHandler(story)

	fmt.Printf("Starting the app on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
