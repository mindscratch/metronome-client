package main

import (
	"fmt"
	"log"

	"github.com/mindscratch/metronome-client/examples/metronome"
	"github.com/mindscratch/metronome-client/api/v1"
)

func main() {
	// disable the schedule
	scheduleJson := `
	{
	  "id": "everyminute",
	  "cron": "* * * * *",
	  "concurrencyPolicy": "ALLOW",
	  "enabled": false,
	  "startingDeadlineSeconds": 60,
	  "timezone": "America/Chicago"
	}
	`

	cl := metronome.Client
	_, err := v1.UpdateSchedule(cl, "prod.example.app", scheduleJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully updated schedule")
}