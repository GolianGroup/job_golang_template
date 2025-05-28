package app

import (
	"jobs_golang_template/internal/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func (a *application) InitMetrics() *metrics.Metrics {
	return metrics.NewMetrics()
}

func (a *application) InitPrometheusRegistry() *prometheus.Registry {
	return prometheus.NewRegistry()

}
