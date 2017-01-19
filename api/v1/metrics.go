// Package v1 provides methods for interacting with the Metronome API v1
package v1

import (
	"net/http"
	"errors"
	"encoding/json"

	"github.com/mindscratch/metronome-client"
	"github.com/mindscratch/metronome-client/types/v1"
)

// Return all of the metronome metrics.
func Metrics(cl client.Client) (*v1.Metrics, error) {

	// Get metrics
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/v1/metrics", nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return nil, errors.New("failed to fetch metrics due to " + err.Error())
	}

	// Parse metrics
	var metrics v1.Metrics
	if err = json.Unmarshal(res, &metrics); err != nil {
		return nil, errors.New("failed to unmarshal JSON data due to " + err.Error())
	}

	return &metrics, nil
}