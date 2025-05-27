package app

import (
	"fmt"
	"jobs_golang_template/internal/config"
	"jobs_golang_template/internal/helper"
	"jobs_golang_template/internal/logging"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func (a *application) InitLogger() (logging.Logger, error) {
	if a.config.Environment == "production" {
		// Config rotation
		ws := zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   a.config.Logger.Rotation.Filename,
				MaxSize:    a.config.Logger.Rotation.MaxSize,
				MaxBackups: a.config.Logger.Rotation.MaxBackups,
				MaxAge:     a.config.Logger.Rotation.MaxAge,
			},
		)

		fluentSyncer, err := helper.NewFluentBitWriteSyncer(
			a.config.Logger.Fluentbit.Host,
			a.config.Logger.Fluentbit.Port,
			a.config.Logger.Fluentbit.Tag,
		)
		if err != nil {
			panic("Failed to create Fluent Bit WriteSyncer: " + err.Error())
		}

		// Combine WriteSyncers
		combinedSyncer := zapcore.NewMultiWriteSyncer(
			ws,                            // lumberjack
			zapcore.AddSync(fluentSyncer), //fluentbit
		)

		// Config encoder and syncer
		level, _ := zapcore.ParseLevel(a.config.Logger.Level)
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config.NewLoggerEncoderConfig(&a.config.Logger.EncoderConfig)),
			combinedSyncer,
			level,
		)
		logger := zap.New(core)
		logger.Sync()

		return logging.NewZapLogger(logger), nil

	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize development logger: %w", err)
	}
	return logging.NewZapLogger(logger), nil
}
