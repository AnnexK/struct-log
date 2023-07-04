package main

import (
	"errors"

	"github.com/AnnexK/struct-log/internal/config"
	rzimpl "github.com/AnnexK/struct-log/internal/infrastructure/logging/rz"
	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
)


func main() {
	cfg := config.MustLoadConfig()
	_ = rzimpl.InitLogging(cfg)

	// Basic logging
	log.Info("hello rz", rz.Int("someField", 64), rz.String("someOtherField", "kek"))

	// Errors
	log.Error("error for rz", rz.Err(errors.New("something bad happened")))

	// Logger context
	logger := log.With(rz.Fields())
}
