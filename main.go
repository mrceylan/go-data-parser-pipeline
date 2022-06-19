package main

import (
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	file, err := readFile("drug-ndc-0001-of-0001.json")
	if err != nil {
		log.Fatal(err)
	}

	errgroup := errgroup.Group{}
	defer file.Close()
	streamer := Streamer[Drug]{
		file:     file,
		channels: reportFuncWrapper(&errgroup, analyseIngredients, analyseStartYear, analyseLatestExpirationDate),
	}
	err = streamFile(streamer, &errgroup)
	if err != nil {
		log.Fatal(err)
	}

	if err := errgroup.Wait(); err != nil {
		log.Fatal("Error", err)
	}

}
