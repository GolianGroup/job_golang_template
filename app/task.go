package app

import (
	"jobs_golang_template/internal/database/scylla"
	"jobs_golang_template/internal/tasks"

	"go.uber.org/zap"
)

func (a *application) InitTask(scyllaDB scylla.ScyllaDB, logger *zap.Logger) tasks.Task {
	return tasks.NewTaskManager(scyllaDB, logger, &a.config.WorkerPool)
}
