package v1

type Metrics struct {
	Counters map[string]Counter `json:"counters"`
	Gauges map[string]Gauge `json:"gauges"`
	Histograms map[string]Histogram `json:"histogram"`
	Version string `json:"version"`
}

type Counter struct {
	Count int64 `json:"count"`
}

type Gauge struct {
	// presumably the value should be a float64, but sometimes the value for a gauge
	// is an empty array, for example: "jvm.threads.deadlocks":{"value":[]}
	// opened issue: https://github.com/dcos/metronome/issues/117
	Value interface{} `json:"value"`
}

type Histogram struct {
	Count int64 `json:"count"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Mean float64 `json:"mean"`
	P50 float64 `json:"p50"`
	P75 float64 `json:"p75"`
	P95 float64 `json:"p95"`
	P98 float64 `json:"p98"`
	P99 float64 `json:"p99"`
	P999 float64 `json:"p999"`
	StdDev float64 `json:"stddev"`
}

type Meter struct {
	Count int64 `json:"count"`
	M15Rate float64 `json:"m15_rate"`
	M1Rate float64 `json:"m1_rate"`
	M5Rate float64 `json:"m5_rate"`
	MeanRate float64 `json:"mean_rate"`
	Units string `json:"units"`
}

type Timer struct {
	Count int64 `json:"count"`
	DurationUnits string `json:"duration_units"`
	M15Rate float64 `json:"m15_rate"`
	M1Rate float64 `json:"m1_rate"`
	M5Rate float64 `json:"m5_rate"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Mean float64 `json:"mean"`
	MeanRate float64 `json:"mean_rate"`
	P50 float64 `json:"p50"`
	P75 float64 `json:"p75"`
	P95 float64 `json:"p95"`
	P98 float64 `json:"p98"`
	P99 float64 `json:"p99"`
	P999 float64 `json:"p999"`
	RateUnits string `json:"rate_units"`
	StdDev float64 `json:"stddev"`
}