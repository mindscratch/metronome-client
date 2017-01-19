package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	metrics, err := v1.Metrics(cl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", metrics)
}
