package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	jobJson := `
	{
	  "description": "Example Application",
	  "id": "prod.example.app",
	  "labels": {
	    "location": "olympus",
	    "owner": "zeus"
	  },
	  "run": {
	    "cmd": "env | sort",
	    "cpus": 0.1,
	    "mem": 64,
	    "disk": 128,
	    "env": {
	      "MON": "test",
	      "CONNECT": "direct"
	    },
	    "maxLaunchDelay": 3600,
	    "user": "root"
	  }
	}
	`

	cl := metronome.Client
	_, err := v1.UpdateJob(cl, jobJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully updated job")
}