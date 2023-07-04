package rz

import (
	"os"

	"github.com/AnnexK/struct-log/internal/config"
	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
)


func InitLogging(cfg config.Config) error {
	level := rz.InfoLevel
	if cfg.Debug {
		level = rz.DebugLevel
	}

	log.SetLogger(log.With(
		rz.Level(level),
		rz.Writer(os.Stdout),
		rz.Formatter(rz.FormatterCLI()),
	))
	return nil
}
