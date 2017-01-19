package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	_, err := v1.DeleteSchedule(cl, "prod.example.app", "everyminute")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted schedule")
}