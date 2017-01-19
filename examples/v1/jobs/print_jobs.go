package main

import (
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	err := v1.PrintJobs(cl, true)
	if err != nil {
		log.Fatal(err)
	}
}
