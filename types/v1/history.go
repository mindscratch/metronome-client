// Package v1 provides types for the Metronome API v1
package v1

// v1/jobs?embed=history
type JobHistory struct {
	SuccessCount int `json:"successCount"`
	FailureCount int `json:"failureCount"`
	LastSuccessAt string `json:"lastSuccessAt"`
	LastFailureAt string `json:"lastFailureAt"`
	SuccessfulFinishedRuns []RunHistory `json:"successfulFinishedRuns"`
	FailedRuns []RunHistory `json:"failedFinishedRuns"`
}

type RunHistory struct {
	Id string `json:"id"`
	CreatedAt string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
}

// v1/jobs?embed=historySummary
type JobHistorySummary struct {
	SuccessCount int `json:"successCount"`
	FailureCount int `json:"failureCount"`
	LastSuccessAt string `json:"lastSuccessAt"`
	LastFailureAt string `json:"lastFailureAt"`
}
