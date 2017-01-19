// Package client provides an HTTP client for Metronome operations
package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client represents the HTTP client for interacting with Metronome
type Client struct {
	Client   *http.Client
	URL      string
	ProxyURL *url.URL
}

// DoRequest makes a request to Metronome REST API
func (cl Client) DoRequest(req *http.Request) ([]byte, error) {

	// Init a client
	client := cl.Client

	if cl.ProxyURL != nil {
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(cl.ProxyURL)}}
	} else {
		if cl.Client == nil {
			client = &http.Client{}
		}
	}

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read data
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return data, errors.New("bad response: " + fmt.Sprintf("%d", resp.StatusCode))
	}

	return data, nil
}

// Returns the Metronome url
func (cl Client) MetronomeUrl() string {
	return strings.TrimSuffix(cl.URL, "/")
}
