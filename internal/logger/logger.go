package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"lk_sut/internal/config"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
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

	return zapCfg.Build()
}
