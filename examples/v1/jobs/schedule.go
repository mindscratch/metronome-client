package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	schedule, err := v1.Schedule(cl, "prod.example.app","everyminute")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", schedule)
}