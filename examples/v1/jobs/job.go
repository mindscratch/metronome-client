package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	job, err := v1.Job(cl, "prod.example.app",false, false, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", job)
}