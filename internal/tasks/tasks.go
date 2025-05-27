package tasks

import (
	"jobs_golang_template/internal/config"
	"jobs_golang_template/internal/database/scylla"

	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type Task interface {
	Start()
	Stop()
	InitWorkerPool() error
}

type task struct {
	scylla     scylla.ScyllaDB
	logger     *zap.Logger
	workerpool *ants.Pool
	quit       chan struct{}
	configs    *config.WorkerPoolConfig
}

func NewTaskManager(
	scyllaDB scylla.ScyllaDB,
	logger *zap.Logger,
	configs *config.WorkerPoolConfig,
) Task {
	return &task{
		scylla:     scyllaDB,
		logger:     logger.With(zap.String("task", "watched")),
		workerpool: nil,
		configs:    configs,
		quit:       make(chan struct{}),
	}
}

func (t *task) InitWorkerPool() error {
	t.logger.Info("Initializing worker pool powered by (YA ALI)", zap.Int("size", t.configs.WorkerPoolSize))
	pool, err := ants.NewPool(t.configs.WorkerPoolSize)
	if err != nil {
		t.logger.Error("Failed to initialize worker pool", zap.Error(err))
		return err
	}
	t.workerpool = pool
	t.logger.Info("Worker pool initialized successfully")
	return nil
}

func (t *task) Start() {
	t.logger.Info("Starting task manager")
	// go t.watchBackgroundJob()
}

func (t *task) Stop() {
	t.logger.Info("Stopping task manager")
	t.quit <- struct{}{}
	t.workerpool.Release()
	t.logger.Info("Task manager stopped successfully")
}
