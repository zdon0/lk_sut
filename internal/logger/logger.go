package logger

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"lk_sut/internal/config"
)

func NewLogger(cfg *config.Config, lc fx.Lifecycle) (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderCfg.StacktraceKey = ""

	logLevel := zap.NewAtomicLevel()

	if cfg.Debug {
		logLevel.SetLevel(zap.DebugLevel)
	}

	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig = encoderCfg
	zapCfg.Level = logLevel

	res, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			// https://github.com/uber-go/zap/issues/328
			_ = res.Sync()
			return nil
		},
	})

	return res, nil
}
