package app

import (
	"context"
	"jobs_golang_template/internal/config"
	"jobs_golang_template/internal/database/postgres"
	"log"

	"go.uber.org/fx"
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
			a.InitTracerProvider,
		),
		fx.Invoke(func(lc fx.Lifecycle, db postgres.Database) {
			shutdownTracer := a.InitTracer()
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					log.Println("starting postgres")
					return nil
				},
				OnStop: func(ctx context.Context) error {
					if err := shutdownTracer(ctx); err != nil {
						log.Printf("Error shutting down tracer: %v", err) // this should change after logging branch get merged
					}
					log.Println(db.Close())
					return nil
				},
			})
		}),
	)
	app.Run()
}
