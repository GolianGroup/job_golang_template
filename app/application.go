package app

import (
	"context"
	"jobs_golang_template/internal/config"
	"jobs_golang_template/internal/database/postgres"
	"jobs_golang_template/internal/logging"
	"jobs_golang_template/internal/metrics"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Application interface {
	Setup()
}

type application struct {
	ctx    context.Context
	config *config.Config
}

func NewApplication(ctx context.Context, config *config.Config) Application {
	return &application{ctx: ctx, config: config}
}

// bootstrap

func (a *application) Setup() {
	app := fx.New(
		fx.Provide(
			a.InitRedis,
			a.InitDatabase,
			a.InitArangoDB,
			a.InitLogger,
			a.InitMetrics,
			a.InitPrometheusRegistry,
		),
		fx.Invoke(func(lc fx.Lifecycle, db postgres.Database) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					log.Println("starting postgres")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println(db.Close())
					return nil
				},
			})
		}),
		fx.Invoke(func(lc fx.Lifecycle, registry *prometheus.Registry, logger logging.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
					go func() {
						if err := http.ListenAndServe(a.config.Server.Host+":"+a.config.Server.Port, nil); err != nil {
							logger.Error("Failed to start Fiber server", zap.Error(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
		fx.Invoke(func(lc fx.Lifecycle, metricsOBJ *metrics.Metrics, registry *prometheus.Registry) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					log.Println("starting metrics")
					registry.MustRegister(
						metricsOBJ.SampleMetric,
						// register your own metrics here
					)
					go func() {
						// Usage example of metrics ( use in your worker logic )
						for {
							mockTemperature := 30 + (5 * (rand.Float64() - 0.5))
							metricsOBJ.SampleMetric.Set(mockTemperature)
							time.Sleep(2 * time.Second)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println("stopping metrics")
					return nil
				},
			})
		}),
	)
	app.Run()
}
