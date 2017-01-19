package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	_, err := v1.DeleteJob(cl, "prod.example.app")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted job")
}