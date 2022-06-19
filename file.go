package main

import (
	"encoding/json"
	"os"

	"golang.org/x/sync/errgroup"
)

type Streamer[T interface{}] struct {
	file     *os.File
	channels []chan<- T
}

func readFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func streamFile[T interface{}](str Streamer[T], errGroup *errgroup.Group) error {
	dec := json.NewDecoder(str.file)
	_, err := dec.Token()
	if err != nil {
		return err
	}

	errGroup.Go(func() error {
		for dec.More() {
			var entity T
			err := dec.Decode(&entity)
			if err != nil {
				return err
			}
			fanoutStream(str.channels, entity)
		}

		_, err = dec.Token()
		if err != nil {
			return err
		}
		closeChannels(str.channels)
		return nil
	})

	return nil
}

func fanoutStream[T interface{}](channels []chan<- T, data T) {
	for _, channel := range channels {
		channel <- data
	}
}

func closeChannels[T interface{}](channels []chan<- T) {
	for _, channel := range channels {
		close(channel)
	}
}
