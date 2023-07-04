package zerolog

import (
	"github.com/AnnexK/struct-log/internal/config"
	zlog "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogging(cfg config.Config) error {
	level := zlog.InfoLevel
	if cfg.Debug {
		level = zlog.DebugLevel
	}

	zlog.SetGlobalLevel(level)

	log.Logger = zlog.New(zlog.NewConsoleWriter()).With().Timestamp().Logger()
	return nil
}
