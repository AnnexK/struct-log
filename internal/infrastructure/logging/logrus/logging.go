package logrus

import (
	"os"

	"github.com/AnnexK/struct-log/internal/config"
	"github.com/sirupsen/logrus"
)

func InitLogging(cfg config.Config) error {
	level := logrus.InfoLevel
	if cfg.Debug {
		level = logrus.DebugLevel
	}

	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	return nil
}
