package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	cl := metronome.Client
	runs, err := v1.Runs(cl, "prod.example.app")
	if err != nil {
		log.Fatal(err)
	}
	if len(runs) > 0 {
		for _, r := range runs {
			fmt.Printf("%#v\n", r)
		}
	} else {
		fmt.Println("no runs currently exist")
	}
}