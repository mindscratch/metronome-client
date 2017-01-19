// Package v1 provides methods for interacting with the Metronome API v1
package v1

import (
	"bytes"
	"strings"
	"os"
	"net/http"
	"errors"
	"encoding/json"

	"github.com/mindscratch/metronome-client"
	"github.com/mindscratch/metronome-client/types/v1"
)

// Return all of the metronome jobs. Optionally, include active runs, schedules, history and/or history summary
// by setting those options to true.
func Jobs(cl client.Client, includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary bool) ([]v1.Job, error) {

	// create the "embed" query parameter, if required
	embedQueryParameter := buildEmbedQueryParameter(includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary)

	// Get jobs
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs"+embedQueryParameter, nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch jobs due to " + err.Error())
	}

	// Parse jobs
	var jobs []v1.Job
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

// Returns the job associated with the given id. Optionally, include active runs, schedules, history and/or history summary
// by setting those options to true.
func Job(cl client.Client, id string, includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary bool) (*v1.Job, error) {
	if id == "" {
		return nil, errors.New("id for a job must be provided")
	}

	// create the "embed" query parameter, if required
	embedQueryParameter := buildEmbedQueryParameter(includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary)

	// Get the job
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs/"+id+embedQueryParameter, nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch job due to " + err.Error())
	}

	// Parse job
	var job v1.Job
	if err = json.Unmarshal(res, &job); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return &job, nil
}

// Create a new job using the provided JSON.
//
// Returns a boolean (success or not) and an error.
func CreateJob(cl client.Client, jobJson string) (bool, error) {
	return doCreateOrUpdateJob(cl, jobJson, true)
}

// Update an existing job with the provided JSON.
//
// Returns a boolean (success or not) and an error.
func UpdateJob(cl client.Client, jobJson string) (bool, error) {
	return doCreateOrUpdateJob(cl, jobJson, false)
}

// Delete an existing job which has the provided id.
//
// Returns a boolean (success or not) and an error.
func DeleteJob(cl client.Client, id string) (bool, error) {
	// check job id
	if id == "" {
		return false, errors.New("id for a job must be provided")
	}

	//  delete the job
	req, err := http.NewRequest("DELETE", cl.MetronomeUrl()+"/v1/jobs/"+id, nil)
	data, err := cl.DoRequest(req)
	if err != nil {
		if data != nil && len(data) > 0 {
			var metronomeError v1.Error
			if unmarshalErr := json.Unmarshal(data, &metronomeError); unmarshalErr != nil {
				return false, errors.New("failed to delete job due to " + err.Error() + " (unable to unmarshal response: " + string(data) + ")")
			} else {
				return false, errors.New("failed to delete job due to " + err.Error() + ": " + string(data))
			}
		} else {
			return false, errors.New("failed to delete job due to " + err.Error())
		}
	}

	return true, nil
}

// Create or update a job using the given JSON. If create is true, the job will be created, otherwise the job will
// be updated.
//
// Returns a boolean (success or not) and an error.
func doCreateOrUpdateJob(cl client.Client, jobJson string, create bool) (bool, error) {
	// check the job, return an error if it's invalid
	buf := []byte(jobJson)
	var job v1.Job
	if err := json.Unmarshal(buf, &job); err != nil {
		return false, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	// create job
	var req *http.Request
	var err error
	if create {
		req, err = http.NewRequest("POST", cl.MetronomeUrl()+"/v1/jobs", bytes.NewBuffer(buf))
	} else {
		req, err = http.NewRequest("PUT", cl.MetronomeUrl()+"/v1/jobs/"+job.Id, bytes.NewBuffer(buf))
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.DoRequest(req)
	if err != nil {
		action := "update"
		if create {
			action = "create"
		}
		return false, errors.New("failed to " + action + " job due to " + err.Error())
	}

	return true, nil
}

func buildEmbedQueryParameter(includeActiveRuns, includeSchedules, includeHistory, includeHistorySummary bool) string {
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
	return embedQueryParameter
}
