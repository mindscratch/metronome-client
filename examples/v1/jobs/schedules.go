package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	jobId := "prod.example.app"
	schedules, err := v1.Schedules(cl, jobId)
	if err != nil {
		log.Fatal(err)
	}
	if len(schedules) > 0 {
		for _, s := range schedules {
			fmt.Printf("%#v\n", s)
		}
	} else {
		fmt.Println("found no schedules for job ", jobId)
	}
}