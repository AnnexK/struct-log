package infrastructure

import "github.com/getsentry/sentry-go"
import "github.com/AnnexK/struct-log/internal/config"

// InitSentry инициализирует Sentry.
func InitSentry(cfg config.Config) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         cfg.Sentry.DSN,
		Release:     cfg.Sentry.Release,
		Environment: cfg.Sentry.Environment,
	}); err != nil {
		return err
	}
	return nil
}
