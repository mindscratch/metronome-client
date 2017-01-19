package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	_, err := v1.StopRun(cl, "prod.example.app","20170119180720D4gdl")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully stopped the run")
}