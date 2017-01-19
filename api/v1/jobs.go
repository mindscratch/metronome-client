// Package v1 provides methods for interacting with the Metronome API v1
package v1

import (
	"os"
	"net/http"
	"errors"
	"encoding/json"
	"fmt"

	"github.com/mindscratch/metronome-client"
	"github.com/mindscratch/metronome-client/types/v1"
	"bytes"
	"strings"
)

// Return all of the metronome jobs. Optionally, include active runs, schedules, history and/or hisotry summary
// by setting those options to true.
func Jobs(cl client.Client, includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary bool) ([]v1.Job, error) {
	embedOptions := make([]string, 0)
	if includeActiveRuns {
		embedOptions = append(embedOptions, "activeRuns")
	}
	if includeSchedules {
		embedOptions = append(embedOptions, "schedules")
	}
	if includeHistory {
		embedOptions = append(embedOptions, "history")
	}
	if includeHistorySummary {
		embedOptions = append(embedOptions, "historySummary")
	}
	embedQueryParameter := ""
	if len(embedOptions) > 0 {
		embedQueryParameter = "embed=" + strings.Join(embedOptions, ",")
	}
	if len(embedQueryParameter) > 0 {
		embedQueryParameter = "?" + embedQueryParameter
	}

	// Get jobs
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs" + embedQueryParameter, nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch jobs due to " + err.Error())
	}

	// Parse jobs
	var jobs []v1.Job
	fmt.Println("REST", string(res))
	if err = json.Unmarshal(res, &jobs); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return jobs, nil
}

// Print the jobs to stdout. The output can optionally be indented by
// setting `pretty` to true.
func PrintJobs(cl client.Client, pretty bool) error {

	// Get jobs
	jobs, err := Jobs(cl, true, true, true, true)
	if err != nil {
		return err
	}

	// Parse jobs
	var buf []byte

	// If pretty is true then
	if pretty {
		buf, err = json.MarshalIndent(jobs, "", "  ")
	} else {
		// Otherwise just parse it
		buf, err = json.Marshal(jobs)
	}

	if err != nil {
		return err
	}

	os.Stdout.Write(buf)

	return nil
}


// Create a new job using the provided JSON.
func CreateJob(cl client.Client, jobJson string) (bool, error) {

	// check the job, return an error if it's invalid
	buf := []byte(jobJson)
	var job v1.Job
	if err := json.Unmarshal(buf, &job); err != nil {
		return false, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	// create job
	req, err := http.NewRequest("POST", cl.MetronomeUrl()+"/v1/jobs", bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.DoRequest(req)
	if err != nil {
		return false, errors.New("failed to create job due to " + err.Error())
	}

	return true, nil
}