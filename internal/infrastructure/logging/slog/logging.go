package slog

import log "golang.org/x/exp/slog"
import "github.com/AnnexK/struct-log/internal/config"
import "os"

// InitLogging инициализирует логгер slog.
func InitLogging(cfg config.Config) error {
	level := log.LevelInfo
	if cfg.Debug {
		level = log.LevelDebug
	}
	logger := log.New(log.NewTextHandler(os.Stdout, &log.HandlerOptions{
		Level: level,
	}))
	log.SetDefault(logger)
	return nil
}
