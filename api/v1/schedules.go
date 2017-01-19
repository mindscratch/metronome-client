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
func Schedules(cl client.Client, jobId string) ([]v1.Schedule, error) {

	// Get schedules
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/schedules", nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch schedules due to " + err.Error())
	}

	// Parse schedules
	var schedules []v1.Schedule
	if err = json.Unmarshal(res, &schedules); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return schedules, nil
}

// Returns the schedule associated with the given job and schedule ids.
func Schedule(cl client.Client, jobId, scheduleId string) (*v1.Schedule, error) {
	if jobId == "" {
		return nil, errors.New("id for a job must be provided")
	}
	if scheduleId == "" {
		return nil, errors.New("id for a schedule must be provided")
	}

	// Get the schedule
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/schedules/"+scheduleId, nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch schedule due to " + err.Error())
	}

	// Parse schedule
	var schedule v1.Schedule
	if err = json.Unmarshal(res, &schedule); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return &schedule, nil
}

// Create a new schedule using the provided JSON.
//
// Returns a boolean (success or not) and an error.
func CreateSchedule(cl client.Client, jobId, scheduleJson string) (bool, error) {
	return doCreateOrUpdateSchedule(cl, jobId, scheduleJson, true)
}

// Update an existing schedule with the provided JSON.
//
// Returns a boolean (success or not) and an error.
func UpdateSchedule(cl client.Client, jobId, scheduleJson string) (bool, error) {
	return doCreateOrUpdateSchedule(cl, jobId, scheduleJson, false)
}

// Delete an existing schedule which has the provided id and is associated with the given job.
//
// Returns a boolean (success or not) and an error.
func DeleteSchedule(cl client.Client, jobId, scheduleId string) (bool, error) {
	if jobId == "" {
		return false, errors.New("id for a job must be provided")
	}
	if scheduleId == "" {
		return false, errors.New("id for a schedule must be provided")
	}

	//  delete the schedule
	req, err := http.NewRequest("DELETE", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/schedules/"+scheduleId, nil)
	data, err := cl.DoRequest(req)
	if err != nil {
		if data != nil && len(data) > 0 {
			var metronomeError v1.Error
			if unmarshalErr := json.Unmarshal(data, &metronomeError); unmarshalErr != nil {
				return false, errors.New("failed to delete schedule due to " + err.Error() + " (unable to unmarshal response: " + string(data) + ")")
			} else {
				return false, errors.New("failed to delete schedule due to " + err.Error() + ": " + string(data))
			}
		} else {
			return false, errors.New("failed to delete schedule due to " + err.Error())
		}
	}

	return true, nil
}

// Create or update a schedule using the given JSON. If create is true, the schedule will be created, otherwise the
// schedule will be updated.
//
// Returns a boolean (success or not) and an error.
func doCreateOrUpdateSchedule(cl client.Client, jobId, scheduleJson string, create bool) (bool, error) {
	// check job id
	if jobId == "" {
		return false, errors.New("id for a job must be provided")
	}

	// check the schedule, return an error if it's invalid
	buf := []byte(scheduleJson)
	var schedule v1.Schedule
	if err := json.Unmarshal(buf, &schedule); err != nil {
		return false, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	// create schedule
	var req *http.Request
	var err error
	if create {
		req, err = http.NewRequest("POST", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/schedules", bytes.NewBuffer(buf))
	} else {
		req, err = http.NewRequest("PUT", cl.MetronomeUrl()+"/v1/jobs/"+jobId+"/schedules/"+schedule.Id, bytes.NewBuffer(buf))
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = cl.DoRequest(req)
	if err != nil {
		action := "update"
		if create {
			action = "create"
		}
		return false, errors.New("failed to " + action + " schedule due to " + err.Error())
	}

	return true, nil
}