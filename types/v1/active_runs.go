// Package v1 provides types for the Metronome API v1
package v1

// v1/jobs?embed=activeRuns
type ActiveRun struct {
	Id string `json:"id"`
	JobId string `json:"jobId"`
	Status string `json:"status"`
	CreatedAt string `json:"createdAt"`
	CompletedAt string `json:"completedAt"`
	Tasks []ActiveTask `json:"tasks"`
}

type ActiveTask struct {
	Id string `json:"id"`
	Status string `json:"status"`
	StartedAt string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}
