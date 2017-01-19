// Package v1 provides types for the Metronome API v1
package v1

type Job struct {
	// REF https://github.com/dcos/metronome/blob/master/api/src/main/resources/public/api/v1/schema/jobspec.schema.json
	Id          string `json:"id"`
	Description string `json:"description"`
	Labels      map[string]string `json:"labels"`
	Run Run `json:"run"`

	// these properties can be populated by providing the `embed` query parameter when querying the /v1/jobs endpoints
	ActiveRuns []ActiveRun `json:"activeRuns,omitempty"`
	HistorySummary *JobHistorySummary `json:"historySummary,omitempty"`
	History *JobHistory `json:"history,omitempty"`
	Schedules []Schedule `json:"schedules,omitempty"`
}

type Run struct {
	Args []string `json:"args"`
	Artifacts []Artifact `json:"artifacts"`
	Cmd string `json:"cmd"`
	Cpus float32 `json:"cpus"`
	Disk int `json:"disk"`
	Docker Docker `json:"docker"`
	Env map[string]string `json:"env"`
	MaxLaunchDelay int `json:"maxLaunchDelay"`
	Mem int `json:"mem"`
	Placement Placement `json:"placement"`
	User string `json:"user"`
	Restart Restart `json:"restart"`
	Volumes []Volume `json:"volumes"`
}

type Artifact struct {
	Uri string `json:"uri"`
	Executable bool `json:"executable"`
	Extract bool `json:"extract"`
	Cache bool `json:"cache"`
}

type Docker struct {
	Image string `json:"image"`
}

type Placement struct {
	Constraints []Constraint `json:"constraints"`
}

type Constraint struct {
	Attribute string `json:"attribute"`
	Operator string `json:"operator"`
	Value string `json:"value"`
}

type Restart struct {
	Policy string `json:"policy"`
	ActiveDeadlineSeconds int `json:"activeDeadlineSeconds"`
}

type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath string `json:"hostPath"`
	Mode string `json:"mode"`
}