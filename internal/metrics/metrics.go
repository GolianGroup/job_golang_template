package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	SampleMetric prometheus.Gauge
	// add your own Metrics here
}

func NewMetrics() *Metrics {
	m := &Metrics{
		SampleMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "sample_metrics_for_job",
			Help: "Current Sample temperature of the CPU.",
		}),
		// define additional Metrics here
	}
	return m
}
