// Package v1 provides methods for interacting with the Metronome API v1
package v1

import (
	"net/http"
	"errors"
	"github.com/mindscratch/metronome-client"
)

func Ping(cl client.Client) (bool, error) {
	// Get jobs
	req, err := http.NewRequest("GET", cl.MetronomeUrl()+"/ping", nil)
	res, err := cl.DoRequest(req)
	if err != nil {
		return false, errors.New("failed to ping due to " + err.Error())
	}
	if string(res) == "pong" {
		return true, nil
	}

	return false, errors.New("failed to ping, expected pong, got " + string(res))
}
