package v1

type Job struct {

// REF https://github.com/dcos/metronome/blob/master/api/src/main/resources/public/api/v1/schema/jobspec.schema.json

Id string `json:"id"`
Description string `json:"description"`
Labels []string `json:"labels"`

}

type Run struct {
  Args []string `json:"args"`
} 
