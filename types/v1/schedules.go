// Package v1 provides types for the Metronome API v1
package v1

type Schedule struct {
	// REF https://github.com/dcos/metronome/blob/master/api/src/main/resources/public/api/v1/schema/schedulespec.schema.json
	Id string `json:"id"`
	Cron string `json:"cron"`
	TimeZone string `json:"timeZone"`
	StartingDeadlineSeconds int `json:"startingDeadlineSeconds"`
	ConcurrencyPolicy string `json:"concurrencyPolicy"`
	Enabled bool `json:"enabled"`
	NextRunAt string `json:"nextRunAt"`
}
