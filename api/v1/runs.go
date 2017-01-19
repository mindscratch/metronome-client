package v1

import (
	"github.com/mindscratch/metronome-client"
	"github.com/mindscratch/metronome-client/types/v1"
	"net/http"
	"errors"
	"encoding/json"
	"bytes"
)

// Return all of the metronome schedules for a particular job.
func Runs(cl client.Client, jobId string) ([]v1.ActiveRun, error) {

	// Get schedules
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/runs", nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch schedules due to " + err.Error())
	}

	// Parse runs
	var runs []v1.ActiveRun
	if err = json.Unmarshal(res, &runs); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return runs, nil
}

// Returns the run associated with the given job and run ids.
func Run(cl client.Client, jobId, runId string) (*v1.ActiveRun, error) {
	if jobId == "" {
		return nil, errors.New("id for a job must be provided")
	}
	if runId == "" {
		return nil, errors.New("id for a run must be provided")
	}

	// Get the run
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/runs/"+runId, nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch run due to " + err.Error())
	}

	// Parse run
	var run v1.ActiveRun
	if err = json.Unmarshal(res, &run); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return &run, nil
}

// Create a new run, meaning force the job to run now.
//
// Returns a boolean (success or not) and an error.
func CreateRun(cl client.Client, jobId string) (bool, error) {
	if jobId == "" {
		return false, errors.New("id for a job must be provided")
	}

	// create run
	buf := []byte("{}")
	req, err := http.NewRequest("POST", cl.MetronomeUrl()+"/v1/jobs/" + jobId + "/runs", bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.DoRequest(req)
	if err != nil {
		return false, errors.New("failed to crate run due to " + err.Error())
	}

	return true, nil
}

// Stops a run.
//
// Returns a boolean (success or not) and an error.
func StopRun(cl client.Client, jobId, runId string) (bool, error) {
	if jobId == "" {
		return false, errors.New("id for a job must be provided")
	}
	if runId == "" {
		return false, errors.New("id for a run must be provided")
	}

	//  stop the run
	req, err := http.NewRequest("POST", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/runs/"+runId+"/action/stop", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")
	data, err := cl.DoRequest(req)
	if err != nil {
		if data != nil && len(data) > 0 {
			var metronomeError v1.Error
			if unmarshalErr := json.Unmarshal(data, &metronomeError); unmarshalErr != nil {
				return false, errors.New("failed to stop run due to " + err.Error() + " (unable to unmarshal response: " + string(data) + ")")
			} else {
				return false, errors.New("failed to stop run due to " + err.Error() + ": " + string(data))
			}
		} else {
			return false, errors.New("failed to stop run due to " + err.Error())
		}
	}

	return true, nil
}
