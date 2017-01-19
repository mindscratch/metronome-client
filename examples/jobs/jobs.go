package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	jobs, err := v1.Jobs(cl, false, false, false, false)
	if err != nil {
		log.Fatal(err)
	}
	for _, j := range jobs {
		//fmt.Printf("%s %s %v\n", j.Id, j.Description)
		fmt.Printf("%#v\n", j)
	}
}