package metronome

import (
	"os"

	"github.com/mindscratch/metronome-client"
)

var (
	// URL is the Metronome url
	URL string

	// Client is the Metronome client
	Client client.Client
)

func init() {
	// Init the Metronome client
	URL = "http://localhost:4400"
	if os.Getenv("METRONOME_URL") != "" {
		URL = os.Getenv("METRONOME_URL")
	}
	Client = client.Client{URL: URL}
}