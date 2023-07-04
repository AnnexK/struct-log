package zap

import (
	"os"

	"github.com/AnnexK/struct-log/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogging(cfg config.Config) error {
	zapcore.AddSync(os.Stdout)
	level := zap.InfoLevel
	if cfg.Debug {
		level = zap.DebugLevel
	}
	loggingConfig := zap.Config{
		Level: zap.NewAtomicLevelAt(level),
		DisableStacktrace: true,
		Encoding: "console",
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := loggingConfig.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	return nil
}
