package v1

type Error struct {
	// REF: github.com/dcos/metronome/api/src/main/scala/dcos/metronome/api/ErrorDetail.scala
	Message string `json:"message"`
}
